from dotenv import load_dotenv
import os

load_dotenv() 

# MongoDB configuration
MONGO_URL = os.getenv("MONGO_URL")
DATABASE_NAME = os.getenv("DATABASE_NAME", "reservationdb")

# RabbitMQ configuration
RABBITMQ_URL = os.getenv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
RABBITMQ_QUEUE = os.getenv("RABBITMQ_QUEUE", "reservation_notifications")

# gRPC server port
GRPC_SERVER_PORT = os.getenv("GRPC_SERVER_PORT", "50051")
