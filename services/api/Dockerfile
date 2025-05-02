# Build stage
FROM golang:1.24 AS builder
WORKDIR /workspace

# Copy go mod files first for better caching
COPY services/api/go.mod services/api/go.sum /workspace/services/api/

# Copy the cate_utils package
COPY libraries/cate_utils /workspace/libraries/cate_utils

WORKDIR /workspace/services/api
RUN go mod download

# Copy source code
COPY services/api ./

# Build the application
RUN go build -o dist/api ./src

RUN chmod +x dist/api

# Final stage
FROM golang:1.24

# Copy only the binary from builder stage
COPY --from=builder /workspace/services/api/dist/api ./api
CMD ["./api"]
