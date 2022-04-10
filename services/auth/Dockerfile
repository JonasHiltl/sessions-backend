FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/cosmtrek/air

COPY ./services/auth ./services/auth
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/auth

RUN go build -o ./auth-service

ENTRYPOINT auth-service
