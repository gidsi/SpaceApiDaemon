FROM golang:alpine

RUN apk add --no-cache git

RUN mkdir /go/src/SpaceApiDaemon

WORKDIR /go/src/app
COPY main .

RUN go-wrapper download
RUN go-wrapper install

ENV SOURCE_COMMIT ${SOURCE_COMMIT:-unknown}

CMD ["go-wrapper", "run"]