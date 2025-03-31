FROM golang:1.24 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcpgo ./cmd/mcpgo

# Use a small alpine image for the final container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/mcpgo .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./mcpgo"]