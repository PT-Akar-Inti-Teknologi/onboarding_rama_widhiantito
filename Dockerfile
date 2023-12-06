# Use the official Golang image as the base image
FROM golang:1.21.4 as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o order_transaction main.go

# Start a new stage
FROM golang:1.21.4

# Set the working directory for the runtime image
WORKDIR /app

# Copy the compiled binary from the builder image
COPY --from=builder /app/order_transaction .

# Expose the port your application runs on
EXPOSE 8080

# Command to run the application
CMD ["./order_transaction"]
