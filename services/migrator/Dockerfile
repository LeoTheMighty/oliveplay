# Use Golang base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Download and install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy migration files
COPY ./migrations /migrations

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o migrate .

# Command to run the migrations
ENTRYPOINT ["./migrate"]