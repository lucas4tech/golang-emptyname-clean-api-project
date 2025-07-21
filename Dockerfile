FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app-challenge ./cmd/main.go

FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache sqlite

COPY --from=builder /app/app-challenge .
COPY --from=builder /app/data ./data

EXPOSE 8080

CMD ["./app-challenge"]