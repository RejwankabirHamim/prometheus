# Stage 1: Build the Go application
FROM golang:latest AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of your application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prom-demo .

# Stage 2: Set up Prometheus and your application
FROM alpine:latest

# Install Prometheus
RUN apk add --no-cache prometheus

# Copy the built Go application from the builder stage
COPY --from=builder /app/prom-demo /app/prom-demo

# Copy your Prometheus configuration file
COPY prometheus.yml /etc/prometheus/prometheus.yml

# Expose ports for both your application and Prometheus
EXPOSE 8181 9090

# Run both your application and Prometheus
CMD ["/bin/sh", "-c", "/app/prom-demo start & /usr/bin/prometheus --config.file=/etc/prometheus/prometheus.yml"]
