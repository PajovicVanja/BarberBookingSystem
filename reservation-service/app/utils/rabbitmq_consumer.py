import json
import pika
import threading
import asyncio
from app.config import RABBITMQ_URL
from app.utils.logger import logger
from app import crud
from app.models import ReservationUpdate

# Name of the queue for payment confirmations
PAYMENT_CONFIRMATIONS_QUEUE = "payment_confirmations"

def process_message(ch, method, properties, body):
    logger.info("Received message: %s", body)
    try:
        data = json.loads(body)
        reservation_id = data.get("reservation_id")
        status = data.get("status")
        if reservation_id and status:
            # Update reservation asynchronously
            loop = asyncio.new_event_loop()
            asyncio.set_event_loop(loop)
            update = ReservationUpdate(status=status)
            loop.run_until_complete(crud.update_reservation(reservation_id, update))
            logger.info("Updated reservation %s to status %s", reservation_id, status)
            loop.close()
        # Acknowledge that message was processed
        ch.basic_ack(delivery_tag=method.delivery_tag)
    except Exception as e:
        logger.error("Error processing message: %s", e)
        # Negative acknowledgement so the message can be requeued
        ch.basic_nack(delivery_tag=method.delivery_tag, requeue=True)

def start_consumer():
    """Set up connection, channel, and start consuming messages."""
    try:
        connection = pika.BlockingConnection(pika.URLParameters(RABBITMQ_URL))
        channel = connection.channel()
        # Declare the queue (durable so that messages survive a broker restart)
        channel.queue_declare(queue=PAYMENT_CONFIRMATIONS_QUEUE, durable=True)
        # Fair dispatch: don’t give more than one message to a worker at a time
        channel.basic_qos(prefetch_count=1)
        channel.basic_consume(queue=PAYMENT_CONFIRMATIONS_QUEUE, on_message_callback=process_message)
        logger.info("Starting RabbitMQ consumer for payment confirmations on queue '%s'", PAYMENT_CONFIRMATIONS_QUEUE)
        channel.start_consuming()
    except Exception as e:
        logger.error("Consumer error: %s", e)
        # Here you could add reconnection logic (e.g., wait and then restart the consumer)
        # For now, we just log the error.

def run_consumer_in_background():
    thread = threading.Thread(target=start_consumer, daemon=True)
    thread.start()
