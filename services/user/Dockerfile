FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/cosmtrek/air

COPY ./services/user ./services/user
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/user

RUN go build -o ./user-service

ENTRYPOINT user-service
