FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o  bin/backend .

EXPOSE 3000

CMD ["./bin/backend"]