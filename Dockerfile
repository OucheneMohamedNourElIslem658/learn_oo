# Use the correct Golang image as the base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o backend .

# Use a minimal Alpine image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/backend .

# Copy the .env file
COPY .env .

# Copy the views directory
COPY services/users/views/ ./services/users/views/

# Expose the port your application runs on
EXPOSE 8080

# Command to run the application
CMD ["./backend"]