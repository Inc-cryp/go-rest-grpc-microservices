## 🏗 Architecture Overview

Client (curl / Postman)
        |
        v
Order Service (REST :8081)
        |
        v
User Service (gRPC :50051)

- Order Service calls User Service via gRPC
- Services communicate using Docker internal DNS
- Configurable via environment variables

---

# 📦 Services

## User Service (gRPC)

- Runs on port 50051
- Provides:
  - GetUserByID
- Returns mock user data (can be extended to PostgreSQL)

---

## Order Service (REST + gRPC Client)

- Runs on port 8081
- Endpoint:

POST /orders/{userID}

Flow:
1. Receive HTTP request
2. Call User Service via gRPC
3. Validate user
4. Create order response

---

# ⚙️ Environment Configuration

Environment variables:

USER_GRPC_ADDR
ORDER_HTTP_PORT
USER_GRPC_PORT
APP_ENV
LOG_LEVEL

---

## 🖥 Local Development (Without Docker)

Create `.env`

USER_GRPC_ADDR=localhost:50051
ORDER_HTTP_PORT=8081

Run services manually:

go run user/main.go
go run order/main.go

---

## 🐳 Run With Docker

Docker uses internal DNS:

USER_GRPC_ADDR=user-service:50051

Example docker-compose.yml:

version: "3.9"

services:
  user-service:
    build:
      context: .
      dockerfile: Dockerfile.user
    ports:
      - "50051:50051"

  order-service:
    build:
      context: .
      dockerfile: Dockerfile.order
    ports:
      - "8081:8081"
    depends_on:
      - user-service
    environment:
      - USER_GRPC_ADDR=user-service:50051
      - ORDER_HTTP_PORT=8081

---

Run Docker:

docker compose up --build

(Note: use "docker compose" not "docker-compose" on latest Docker)

---

# Test API

Create order:

curl -X POST http://localhost:8081/orders/1

If user exists → order created  
If not → returns:

{"error":"user not found"}

---

# Scalability Plan

This project can be scaled into production ready architecture:

### PostgreSQL Integration
- Store users & orders persistently
- Replace mock data with DB queries

### Redis Integration
- Cache user data
- Improve performance
- Reduce gRPC calls

### Event-Driven (Kafka / NATS)
- Publish OrderCreated events
- Async processing

### Improvements
- gRPC retry mechanism
- Health checks
- Logging middleware
- Graceful shutdown
- Centralized config
- CI/CD pipeline

#  Why This Project Matters

- Microservice communication (REST → gRPC)
- Docker networking
- Environment-based configuration
- Service dependency handling
- Clean service separation

