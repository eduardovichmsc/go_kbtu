# Use a newer version of the Go base image, e.g., 1.25.6 or the latest stable version
FROM golang:1.25.6 AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build "-o" practice_4 ./cmd/api/main.go

# (Optional) Second stage for a minimal runtime image
# FROM scratch AS final
# ... copy binary from builder stage
