# app/utils/rabbitmq_consumer.py
import json
import pika
import threading
import asyncio
from bson.errors import InvalidId
from app.config import RABBITMQ_URL
from app.utils.logger import logger
from app import crud
from app.models import ReservationUpdate
from app.utils.publisher import publish_event

# Name of the queue for payment confirmations
PAYMENT_CONFIRMATIONS_QUEUE = "payment_confirmations"

def process_message(ch, method, properties, body):
    logger.info("Received message: %s", body)
    try:
        data = json.loads(body)
        reservation_id = data.get("reservation_id")
        status = data.get("status")
        if reservation_id and status:
            # Attempt to update the reservation in MongoDB
            loop = asyncio.new_event_loop()
            asyncio.set_event_loop(loop)
            update = ReservationUpdate(status=status)
            try:
                loop.run_until_complete(crud.update_reservation(reservation_id, update))
                logger.info("Updated reservation %s to status %s", reservation_id, status)
                # Publish domain event for status change
                publish_event("ReservationStatusUpdated", {
                    "reservation_id": reservation_id,
                    "status": status
                })
            except InvalidId:
                # Invalid ObjectId: log and ack so it won’t requeue
                logger.error("Invalid reservation ID format: %s", reservation_id)
            finally:
                loop.close()
        # Acknowledge *all* messages here so none ever requeue endlessly
        ch.basic_ack(delivery_tag=method.delivery_tag)

    except json.JSONDecodeError as e:
        logger.error("Malformed JSON, dropping message: %s", e)
        ch.basic_ack(delivery_tag=method.delivery_tag)

    except Exception as e:
        # For any other unexpected error, ack to avoid poison‐message loops.
        logger.error("Error processing message, dropping: %s", e)
        ch.basic_ack(delivery_tag=method.delivery_tag)


def start_consumer():
    """Set up connection, channel, and start consuming messages."""
    while True:
        try:
            connection = pika.BlockingConnection(pika.URLParameters(RABBITMQ_URL))
            channel = connection.channel()
            # Declare the queue (durable)
            channel.queue_declare(queue=PAYMENT_CONFIRMATIONS_QUEUE, durable=True)
            channel.basic_qos(prefetch_count=1)
            channel.basic_consume(
                queue=PAYMENT_CONFIRMATIONS_QUEUE,
                on_message_callback=process_message
            )
            logger.info(
                "Starting RabbitMQ consumer for payment confirmations on queue '%s'",
                PAYMENT_CONFIRMATIONS_QUEUE
            )
            channel.start_consuming()
        except Exception as e:
            logger.error("Consumer error, retrying in 5s: %s", e)
            try:
                channel.close()
                connection.close()
            except:
                pass
            time.sleep(5)


def run_consumer_in_background():
    thread = threading.Thread(target=start_consumer, daemon=True)
    thread.start()
