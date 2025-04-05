FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git and other dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Use a minimal alpine image for the final container
FROM alpine:3.18

WORKDIR /app

# Install dependencies
RUN apk add --no-cache ca-certificates

# Copy the binary from builder
COPY --from=builder /app/server .

# Create a non-root user to run the app
RUN adduser -D -g '' appuser
USER appuser

# Expose the API port
EXPOSE 8080

# Command to run
CMD ["./server"]