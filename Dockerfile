FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod tidy && go mod vendor

RUN go build -o  bin/backend .

EXPOSE 3000

CMD ["./bin/backend"]