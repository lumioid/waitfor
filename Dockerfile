#
# Lumotion Worker44 Daemons
#
FROM golang:1.16-alpine
MAINTAINER Erwin Saputra <erwin@lumio.id>
COPY . /go/src/github.com/lumio/waitfor/

# Build the binary first
RUN cd /go/src/github.com/lumio/waitfor/cmd/waitfor/ \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build \
      -a \
      -mod vendor \
      -installsuffix cgo \
      -o "/go/bin/waitfor" ./main.go

# multistage build
FROM alpine:3.13
RUN addgroup -g 1000 waitfor \
    && adduser -u 1000 -G waitfor -s /bin/sh -D waitfor \
    && mkdir -p /srv \
    && chown waitfor:waitfor /srv

COPY --from=0 /go/bin/waitfor /srv/waitfor

USER waitfor

CMD ["/srv/waitfor"]
