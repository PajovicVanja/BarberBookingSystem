import pika
import json
from app.config import RABBITMQ_URL, RABBITMQ_EVENT_QUEUE
from app.utils.logger import logger

def publish_event(event_type: str, data: dict):
    """
    Publish a domain event to RabbitMQ.
    """
    try:
        connection = pika.BlockingConnection(pika.URLParameters(RABBITMQ_URL))
        channel = connection.channel()
        channel.queue_declare(queue=RABBITMQ_EVENT_QUEUE, durable=True)

        message = {
            "type": event_type,
            "data": data
        }
        channel.basic_publish(
            exchange='',
            routing_key=RABBITMQ_EVENT_QUEUE,
            body=json.dumps(message),
            properties=pika.BasicProperties(delivery_mode=2)
        )
        logger.info(f"Published event {event_type}: {data}")
        connection.close()
    except Exception as e:
        logger.error(f"Failed to publish event {event_type}: {e}")
