FROM golang:1.20.3-alpine as local

RUN apk update && apk add git && apk add bash && apk add vim

ARG WORKDIR=/go/app/src
WORKDIR $WORKDIR

COPY go.mod $WORKDIR
COPY .air.toml $WORKDIR

RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4/middleware
RUN go get github.com/cosmtrek/air@latest

CMD go mod tidy && air -c .air.toml