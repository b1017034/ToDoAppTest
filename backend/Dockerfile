FROM golang:1.20-alpine

ENV GOPATH /go
ENV GO111MODULE on

RUN apk update && apk add git

WORKDIR /go/src
COPY go.mod .

RUN go mod tidy && \
    go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]