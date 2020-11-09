FROM golang:1.13.5-alpine3.10 AS build-env
WORKDIR /go/src/github.com/miraikeitai2020/backend-log
COPY ./ ./
RUN apk add --no-cache git
RUN go build -o server ../cmd/main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-log/build/server /usr/local/bin/server
ENV PORT 8080

EXPOSE 8080
CMD ["/usr/local/bin/server"]