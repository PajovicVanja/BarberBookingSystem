# sendtest.py
import pika
import json

connection = pika.BlockingConnection(pika.URLParameters("amqp://guest:guest@localhost:5672/"))
channel = connection.channel()
channel.queue_declare(queue="reservation_notifications", durable=True)

message = {
    "reservation_id": "12345",
    "status": "confirmed"
}
channel.basic_publish(
    exchange='',
    routing_key='reservation_notifications',
    body=json.dumps(message),
    properties=pika.BasicProperties(delivery_mode=2)
)
print("✅ Sent test message")
connection.close()
