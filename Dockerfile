# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Run the built binary in a small container
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy static files if needed
COPY public/ public/

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["./main"]
