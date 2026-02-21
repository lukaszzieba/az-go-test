# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s' \
    -o bin/go-fun \
    ./cmd/api/main.go

# Runtime stage - scratch (ultra-minimal)
FROM scratch

# Copy binary from builder
COPY --from=builder /app/bin/go-fun /go-fun

# Expose port
EXPOSE 8081

# Run application
ENTRYPOINT ["/go-fun"]
