# Use the official Go image as a parent image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy Go source files
COPY ./app ./app
COPY ./database ./database
COPY ./docs ./docs
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Build the Go application
RUN go build -o main ./app/bootstrap

# Expose the port the app runs on
EXPOSE 8070

# Command to run when the container starts
CMD ["./main"]
