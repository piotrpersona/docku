FROM golang:1.12.1-alpine3.9 AS stage-build

WORKDIR /go/src/github.com/piotrpersona/docku

RUN apk update && apk add bash git \
    && go get -u github.com/kardianos/govendor

COPY vendor/vendor.json vendor/vendor.json
RUN govendor sync

COPY . .

ARG distro="linux/amd64"
ARG OUTPUT_NAME="docku"
RUN ./ci/build.sh "${distro}"

FROM docker:18.06.1

COPY --from=stage-build \
    /go/bin/docku /usr/local/bin/docku

ENTRYPOINT [ "/usr/local/bin/docku" ]
