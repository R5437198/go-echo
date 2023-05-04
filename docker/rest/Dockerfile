FROM golang:1.20.0-alpine as local

RUN apk update && apk add git && apk add bash && apk add vim && apk add make

ARG WORKDIR=/go/app/src
WORKDIR $WORKDIR

COPY go.mod $WORKDIR
COPY .air.toml $WORKDIR

RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4/middleware
RUN go install -tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
RUN go install github.com/cosmtrek/air@latest

CMD go mod tidy && air -c .air.toml