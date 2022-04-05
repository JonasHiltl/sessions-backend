FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/cosmtrek/air

COPY ./services/profile ./services/profile
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/profile

RUN ls -la

RUN go build -o ./profile-service

RUN ls -la

ENTRYPOINT profile-service
