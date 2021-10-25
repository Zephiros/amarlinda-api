FROM golang:alpine

WORKDIR /amarlinda

ADD . .

RUN go mod download

ENTRYPOINT go build && ./amarlinda

EXPOSE 8082
