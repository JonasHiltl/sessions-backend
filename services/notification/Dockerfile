FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
RUN go get github.com/cosmtrek/air

COPY . ./

EXPOSE 8080

WORKDIR /app/services/notification

RUN go build -o /notification-service

CMD [ "/notification-service" ]
