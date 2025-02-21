# Use latest Go version
FROM golang:1.24 AS builder

WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum /app/
RUN go mod download

# Copy all source code
COPY . /app/

# Build the application
RUN go build -o food-delivery-app .

# Create lightweight runtime environment
FROM ubuntu:latest

WORKDIR /root/

# Copy built binary
COPY --from=builder /app/food-delivery-app /root/

# Copy .env file
COPY --from=builder /app/.env /root/

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./food-delivery-app"]
