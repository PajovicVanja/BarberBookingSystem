FROM golang:1.24-alpine

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 🔍 Debug: print all files in /app after copy
RUN echo "📂 Listing all files in /app:" && ls -R /app

# 🔨 Build binary
RUN go build -o paymentservice ./cmd/payment-service && echo "✅ Build succeeded" && ls -lh paymentservice

EXPOSE 8080

CMD ["./paymentservice"]
