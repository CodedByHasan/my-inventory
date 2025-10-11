# Start from the official Golang base image
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Copy my-inventory source code
COPY go.mod go.sum ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY utils/ ./utils/

RUN go mod download && \
    go build -o api-server ./cmd/api-server/main.go

# Use a minimal image for the final container
FROM debian:bookworm-slim

# Set working directory in the minimal image
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api-server .

ARG APP_PORT
ENV APP_PORT=${APP_PORT}
# Expose the port your app runs on (customize as needed)
EXPOSE ${APP_PORT}

# Run the binary
CMD ["./api-server"]
