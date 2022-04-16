FROM golang:1.18.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/cosmtrek/air

COPY ./services/auth ./services/auth
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/auth

RUN go build -o ./auth-service

CMD ["./auth-aggregator"]
