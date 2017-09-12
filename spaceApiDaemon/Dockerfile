FROM golang:alpine

ARG SOURCE_COMMIT=unknown
ARG DOCKER_TAG=unknown

RUN apk add --no-cache git

RUN mkdir /go/src/SpaceApiDaemon

WORKDIR /go/src/app
COPY main .

RUN go-wrapper download
RUN go-wrapper install

ENV SOURCE_COMMIT=${SOURCE_COMMIT}
ENV DOCKER_TAG=${DOCKER_TAG}

CMD ["go-wrapper", "run"]