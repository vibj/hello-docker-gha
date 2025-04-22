# Use the official Golang image to build the application
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o /helloworld

# Use a minimal image for the final container
FROM alpine:latest

# Copy the compiled binary from the builder stage
COPY --from=builder /helloworld /helloworld

# Set the entrypoint to the binary
ENTRYPOINT ["/helloworld"]