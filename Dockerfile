# Use the official Golang image
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp cmd/main.go

EXPOSE 8080

CMD ["./myapp"]
