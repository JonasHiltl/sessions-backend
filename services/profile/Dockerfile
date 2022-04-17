FROM golang:1.18.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY ./services/profile ./services/profile
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/profile

RUN go build -o profile-service

ENTRYPOINT ./profile-service
