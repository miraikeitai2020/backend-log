FROM golang:1.13.5-alpine3.10 AS build-env
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/miraikeitai2020/backend-log
COPY build ./
RUN apk add --no-cache git
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-log/server /usr/local/bin/server

CMD /usr/local/bin/server $PORT