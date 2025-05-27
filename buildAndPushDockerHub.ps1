# PowerShell script to build, tag, and push all Docker images

Write-Host "🔧 Building and pushing backend services..."

# User Service
docker build -t user-service:latest user-service/
docker tag user-service:latest pajovicv/user-service:latest
docker push pajovicv/user-service:latest

# Reservation Service
docker build -t reservation-service:latest reservation-service/
docker tag reservation-service:latest pajovicv/reservation-service:latest
docker push pajovicv/reservation-service:latest

# Payment Service
docker build -t payment-service:latest payment-service/
docker tag payment-service:latest pajovicv/payment-service:latest
docker push pajovicv/payment-service:latest

# Gateway Web
docker build -t gateway-web:latest gateway-web/
docker tag gateway-web:latest pajovicv/gateway-web:latest
docker push pajovicv/gateway-web:latest

# Gateway Mobile
docker build -t gateway-mobile:latest gateway-mobile/
docker tag gateway-mobile:latest pajovicv/gateway-mobile:latest
docker push pajovicv/gateway-mobile:latest

Write-Host "🎨 Building and pushing frontend microfrontends..."

# Container Frontend (Host App)
docker build -t container-frontend:latest frontend/container/
docker tag container-frontend:latest pajovicv/container-frontend:latest
docker push pajovicv/container-frontend:latest

# Users Microfrontend
docker build -t users-mf:latest frontend/users-mf/
docker tag users-mf:latest pajovicv/users-mf:latest
docker push pajovicv/users-mf:latest

# Reservations Microfrontend
docker build -t reservations-mf:latest frontend/reservations-mf/
docker tag reservations-mf:latest pajovicv/reservations-mf:latest
docker push pajovicv/reservations-mf:latest

# Payments Microfrontend
docker build -t payments-mf:latest frontend/payments-mf/
docker tag payments-mf:latest pajovicv/payments-mf:latest
docker push pajovicv/payments-mf:latest

Write-Host "✅ All Docker images built and pushed successfully."
