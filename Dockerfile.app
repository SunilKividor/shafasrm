FROM golang:1.23-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary ./cmd/app/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app-binary /app/app-binary

EXPOSE 3000

CMD ["/app/app-binary"]