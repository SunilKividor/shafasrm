FROM golang:1.23-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apk add --no-cache make

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

EXPOSE 3000

CMD ["make", "run"]