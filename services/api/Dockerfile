# Use the official Golang image
FROM golang:1.20-alpine

WORKDIR /app
COPY . .

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -o main .

EXPOSE 3000
CMD ["./main"]
