FROM golang:alpine

WORKDIR /amarlinda

ADD . .

RUN go mod download

ENTRYPOINT go build && ./bin/air

EXPOSE 8082
