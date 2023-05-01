FROM golang:1.20-alpine

ARG ORG

ENV GOPATH /go
ENV GO111MODULE on
ENV GOPROXY direct


RUN apk update && apk add git

WORKDIR /go/src
COPY go.mod .

RUN go mod tidy && \
    go env -w GOPRIVATE=github.com/${ORG} && \
    git config --global url."ssh://git@github.com/${ORG}/".insteadOf "https://github.com/${ORG}/" && \
    go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]