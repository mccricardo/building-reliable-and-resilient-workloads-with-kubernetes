# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Stage 2: Create the final image
FROM gcr.io/distroless/static

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Set the entrypoint
ENTRYPOINT ["/app/main"]
