FROM golang:1.12.1-alpine3.9 AS stage-build

WORKDIR /go/src/docku

COPY . .

RUN go build -o /go/bin/docku main.go

FROM docker:18.06.1

COPY --from=stage-build \
    /go/bin/docku /usr/local/bin/docku

ENTRYPOINT [ "/usr/local/bin/docku" ]
