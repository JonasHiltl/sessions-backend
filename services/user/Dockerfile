FROM golang:1.18.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY ./services/user ./services/user
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/user

RUN go build -o user-service

ENTRYPOINT ./user-service
