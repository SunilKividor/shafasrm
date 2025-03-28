FROM golang:1.21-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

EXPOSE 3000

CMD ["make", "run"]