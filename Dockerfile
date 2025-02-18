# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application
COPY . .

# ✅ Generate Swagger documentation inside the container
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init

# Build the application
RUN GOOS=linux GOARCH=amd64 go build -o hermes-app .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy built binary and Swagger docs from builder
COPY --from=builder /app/hermes-app .
COPY --from=builder /app/docs ./docs

# Expose the API port
EXPOSE 8080

# Run the Hermes API
CMD ["./hermes-app"]