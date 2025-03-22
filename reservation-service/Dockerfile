# Use official Python image
FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Copy requirements first
COPY requirements.txt .

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Install supervisord
RUN apt-get update && apt-get install -y supervisor && apt-get clean

# Copy all source files
COPY . .

# Generate gRPC Python files using `--proto_path=app/proto` and `--python_out=.` (not ./app)
RUN python -m grpc_tools.protoc -I./app/proto \
    --python_out=. \
    --grpc_python_out=. \
    app/proto/reservation.proto

# Supervisor config
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Ports
EXPOSE 8000
EXPOSE 50051

# ✅ Set PYTHONPATH to /app so `from app import ...` works in Docker
ENV PYTHONPATH=/app

# Start services
CMD ["/usr/bin/supervisord"]
