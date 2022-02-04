FROM golang:1.17-alpine AS builder

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN  CGO_ENABLED=0 GOOS=linux go build -o usr/bin/user ./services/user
FROM alpine

EXPOSE 8081
WORKDIR /root/

COPY --from=builder work/usr/bin/user /usr/bin/user
CMD ["/usr/bin/user"]