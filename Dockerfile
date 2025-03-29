FROM golang:1.23-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apk add --no-cache make

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN make up

RUN go build -o my-app ./cmd/app/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/my-app .

EXPOSE 3000

CMD ["./my-app"]