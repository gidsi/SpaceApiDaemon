FROM golang:alpine

RUN apk add --no-cache git

RUN mkdir /go/src/SpaceApiDaemon

WORKDIR /go/src/app
COPY main .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]