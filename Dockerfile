FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go mod vendor

RUN go build -C cmd/app -o ../../bin/backend .

EXPOSE 3000

CMD ["./bin/backend"]