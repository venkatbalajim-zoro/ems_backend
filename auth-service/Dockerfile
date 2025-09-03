# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Preload dependencies
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download

# Copy source and build
COPY . .
RUN go build -o app .

# Final stage
FROM alpine:latest
WORKDIR /app

# Copy only the binary from builder
COPY --from=builder /app/app .

# Expose the port Vercel expects
EXPOSE 3000

# Run the service
CMD ["./app"]
