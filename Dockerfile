FROM golang:1.22.7-alpine3.20 AS builder

ENV CGO_ENABLED=0

COPY . .

RUN go mod download

RUN go build -o /tmp/example-wrap-process .


FROM busybox:1.37.0-musl

COPY --from=builder --chown=1000:1000 /tmp/example-wrap-process example-wrap-process

USER 1000:1000

ENTRYPOINT ["cp", "example-wrap-process", "/transfer/"]