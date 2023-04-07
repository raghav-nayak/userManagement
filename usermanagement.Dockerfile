# Start from the latest golang base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the Go mod and sum files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY *.go .

# Build the Go application
RUN go build -o userManagement .

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./userManagement"]
