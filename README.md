# README.md - Documentation for Hermes API

# Hermes API
Hermes is a message queue service using Redis Streams, designed to facilitate secure and scalable pub/sub messaging.

## Features
- Secure API key authentication
- Redis Streams for durable messaging
- API endpoints for publishing and subscribing
- Admin API for managing whitelisted apps
- Swagger documentation for easy testing

## Installation

### Prerequisites
- Docker & Docker Compose
- Go (if running locally without Docker)

### Clone the Repository
```bash
git clone https://github.com/your-repo/hermes.git
cd hermes
```

### Running with Docker Compose
```bash
docker-compose up -d --build
```

### Running Locally
```bash
go mod tidy
go build -o hermes-app .
./hermes-app
```

## API Usage

### Register an App
```http
POST /register
```
Request Body:
```json
{
  "app_name": "my-app"
}
```
Response:
```json
{
  "api_key": "abcd1234efgh5678"
}
```

### Publish a Message
```http
POST /publish
```
Headers:
```http
Authorization: ApiKey abcd1234efgh5678
```
Request Body:
```json
{
  "channel": "hermes_channel",
  "message": "Hello from Hermes!"
}
```
Response:
```json
{
  "status": "message published"
}
```

### Subscribe to Messages
```http
POST /subscribe
```
Headers:
```http
Authorization: ApiKey abcd1234efgh5678
```
Request Body:
```json
{
  "channel": "hermes_channel"
}
```
Response:
```json
{
  "messages": ["Hello from Hermes!"]
}
```

### Manage Whitelisted Apps (Admin)
```http
POST /admin/whitelist
```
Headers:
```http
Authorization: AdminKey supersecurekey
```
Request Body:
```json
{
  "app_name": "new-app",
  "action": "add"
}
```
Response:
```json
{
  "status": "whitelist updated"
}
```

### Check API Health
```http
GET /health
```
Response:
```json
{
  "status": "ok",
  "redis": "connected"
}
```

## Swagger API Documentation
Once the service is running, visit:
```
http://localhost:8080/docs
```
to test the API endpoints.

## License
Copywrite IVBH