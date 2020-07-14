 
FROM golang:1.12.0-alpine3.9

RUN apk update && apk add git && go get gopkg.in/natefinch/lumberjack.v2

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/cmd/catalog/main"]