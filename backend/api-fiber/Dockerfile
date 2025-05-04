# Use the official Go image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Set environment variables
ENV DB_HOST=mysql_goose
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASSWORD=mi_password
ENV DB_NAME=cursos
ENV GOOSE_DRIVER=mysql
ENV GOOSE_DBSTRING="root:mi_password@tcp(mysql_goose:3306)/cursos?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
ENV GOOSE_MIGRATION_DIR=./database/migrations
ENV GOOSE_TABLE=custom.goose_migrations

# Install goose CLI tool
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the Go source files into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8070

# Command to run when the container starts
CMD ["./main"]
