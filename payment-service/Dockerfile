# Start from the official Golang image
FROM golang:1.18-alpine

# Set Go environment variables
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Create and set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application binary
RUN go build -o paymentservice ./cmd/payment-service

# Expose the port (default 8080)
EXPOSE 8080

# Run the binary
CMD ["./paymentservice"]
