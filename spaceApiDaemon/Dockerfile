FROM golang:1.10.2-alpine

ARG SOURCE_COMMIT=unknown
ARG DOCKER_TAG=unknown

RUN apk add --no-cache git

WORKDIR /go/src/app
COPY main .

RUN go get -d -v ./...
RUN go install -v ./...

ENV SOURCE_COMMIT=${SOURCE_COMMIT}
ENV DOCKER_TAG=${DOCKER_TAG}

CMD ["app"]