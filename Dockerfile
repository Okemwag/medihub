# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o medihub ./cmd/medihub

# Use a minimal Alpine image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/medihub .

# Copy the migrations directory
COPY ./migrations ./migrations

# Copy the .env file
COPY .env .

# Copy the Air binary from the builder stage
COPY --from=builder /go/bin/air /usr/local/bin/air

# Expose the application port
EXPOSE 8000

# Command to run the application
CMD ["air", "-c", ".air.toml"]