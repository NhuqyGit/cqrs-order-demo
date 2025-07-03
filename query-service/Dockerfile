# Dockerfile for query-service
FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o query-service

EXPOSE 8081
CMD ["./query-service"]
