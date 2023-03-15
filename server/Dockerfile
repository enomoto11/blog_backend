FROM golang:1.19.2-alpine3.16 AS dev

WORKDIR /go/src
RUN apk update && apk add alpine-sdk # gcc, gitなど入っている
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]