FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/cosmtrek/air

COPY ./services/data_aggregator ./services/data_aggregator
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/data_aggregator

RUN go build -o ./data-aggregator

CMD ["./data-aggregator"]
