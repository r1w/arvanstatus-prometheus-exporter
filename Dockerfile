# Stage 1: Build the Go application
FROM golang:1.23.0-alpine3.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o arvanstatus-exporter

# Stage 2: Create a minimal runtime image
FROM alpine:3.20

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/arvanstatus-exporter .

# Expose the necessary port (replace with your app's port)
EXPOSE 8002

# Run the exporter
CMD ["./arvanstatus-exporter"]
