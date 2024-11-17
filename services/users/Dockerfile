# Use the official Golang image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from the root of the project
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy only the users service code
COPY services/users/ ./

# Build the Go application
RUN go build -o main .

# Expose the port that the service will run on
EXPOSE 8080

# Set the command to run the application
CMD ["./main"]
