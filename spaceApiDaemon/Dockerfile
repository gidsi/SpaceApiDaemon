FROM golang:alpine

RUN apk add --no-cache git

RUN mkdir /go/src/SpaceApiDaemon

WORKDIR /go/src/app
COPY main .

RUN go-wrapper download
RUN go-wrapper install

ARG SOURCE_COMMIT=${SOURCE_COMMIT:-unknown}
ARG DOCKER_TAG=${DOCKER_TAG:-unknown}

RUN echo $SOURCE_COMMIT
RUN echo $DOCKER_TAG

CMD ["go-wrapper", "run"]