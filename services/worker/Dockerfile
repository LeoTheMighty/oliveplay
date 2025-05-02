# Build stage
FROM golang:1.24 AS builder
WORKDIR /workspace

# Copy in go.mod and go.sum first for caching
COPY services/worker/go.mod services/worker/go.sum /workspace/services/worker/

# Copy the cate_utils package
COPY libraries/cate_utils /workspace/libraries/cate_utils

WORKDIR /workspace/services/worker
RUN go mod download

# Copy entire worker source
COPY services/worker ./

# Build the worker
RUN go build -o dist/worker ./src

RUN chmod +x dist/worker

# Final stage
FROM golang:1.24

# Copy only the binary from builder stage
COPY --from=builder /workspace/services/worker/dist/worker ./worker
CMD ["./worker"]
