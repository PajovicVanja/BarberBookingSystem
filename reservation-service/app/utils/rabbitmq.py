import pika
from app.config import RABBITMQ_URL, RABBITMQ_QUEUE
from app.utils.logger import logger

def send_confirmation(reservation_id: str):
    try:
        connection = pika.BlockingConnection(pika.URLParameters(RABBITMQ_URL))
        channel = connection.channel()
        channel.queue_declare(queue=RABBITMQ_QUEUE, durable=True)
        message = f"Reservation {reservation_id} confirmed"
        channel.basic_publish(
            exchange='',
            routing_key=RABBITMQ_QUEUE,
            body=message,
            properties=pika.BasicProperties(delivery_mode=2)
        )
        logger.info(f"Sent confirmation for reservation {reservation_id}")
        connection.close()
    except Exception as e:
        logger.error(f"Failed to send confirmation: {e}")
