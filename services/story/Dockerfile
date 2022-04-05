FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go get github.com/cosmtrek/air

COPY ./services/story ./services/story
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/story

RUN go build -o ./story-service

ENTRYPOINT story-service
