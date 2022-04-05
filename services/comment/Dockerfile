FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go get github.com/cosmtrek/air

COPY ./services/comment ./services/comment
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/comment

RUN go build -o ./comment-service

ENTRYPOINT comment-service
