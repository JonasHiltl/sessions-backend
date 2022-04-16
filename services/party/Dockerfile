FROM golang:1.18.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/cosmtrek/air

COPY ./services/party ./services/party
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/party

RUN go build -o ./party-service

CMD ["./party-service"]
