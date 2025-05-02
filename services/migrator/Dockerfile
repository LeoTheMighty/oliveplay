# Build stage
FROM golang:1.24 AS builder
WORKDIR /workspace

RUN apt-get update && apt-get install -y file

# Copy go.mod, go.sum
COPY services/migrator/go.mod services/migrator/go.sum ./services/migrator/

# Copy the cate_utils package
COPY libraries/utils /workspace/libraries/utils

WORKDIR /workspace/services/migrator
RUN go mod download

# Copy the rest
COPY services/migrator ./

# Build the migrator
RUN go build -o dist/migrator ./src
RUN chmod +x dist/migrator

RUN file ./dist/migrator

# Final stage
FROM golang:1.24

# Copy the binary from builder stage
COPY --from=builder /workspace/services/migrator/dist/migrator ./migrator

# Copy the migrations
COPY --from=builder /workspace/services/migrator/src/migrations ./migrations

CMD ["./migrator"]
