# Dockerfile for cmd-service
FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o cmd-service

EXPOSE 8080
CMD ["./cmd-service"]
