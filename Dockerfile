# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the entire source code from the current directory to the working directory in the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
