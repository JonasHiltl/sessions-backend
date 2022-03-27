FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
COPY ./services ./services
COPY ./packages ./packages
RUN go mod tidy
RUN go get github.com/cosmtrek/air

EXPOSE 8081

WORKDIR /app/services/story

RUN go build -o /story-service

CMD [ "/story-service" ]
