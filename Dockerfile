# Build stage
FROM golang:1.24.3-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o user-service

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user-service .
CMD ["./user-service"]
