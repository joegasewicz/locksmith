# syntax=docker/dockerfile:1
FROM golang:1.21-alpine

WORKDIR /identity_api

COPY . .

RUN go mod download

RUN go build -o /identity_server

EXPOSE 7001

CMD ["/identity_server"]
