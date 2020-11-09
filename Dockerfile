FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/miraikeitai2020/backend-log
COPY . .
RUN go build -o server main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/miraikeitai2020/backend-log/server /app

CMD /app/main $PORT