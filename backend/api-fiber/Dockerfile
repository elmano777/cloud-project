# Build stage
FROM golang:1.21 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Run stage
FROM debian:bullseye-slim

# Instalar dependencias de ejecución (si necesitas algo como libc para correr el binario)
RUN apt-get update && apt-get install -y libc6

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8070
CMD ["./main"]
