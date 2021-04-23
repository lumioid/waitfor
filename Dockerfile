#
# Lumotion Worker44 Daemons
#
FROM golang:1.16-alpine
MAINTAINER Erwin Saputra <erwin@lumio.id>
COPY . /go/src/github.com/lumio/waitfor/

# Build the binary first - we opt in for upx compressed binary since
# this image need to be small and is non persistent
RUN apk add upx
RUN cd /go/src/github.com/lumio/waitfor/cmd/waitfor/ \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build \
      -a \
      -mod vendor \
      -ldflags="-s -w" \
      -o "/go/bin/waitfor" ./main.go \
    && upx /go/bin/waitfor

# multistage build
FROM scratch
COPY --from=0 /go/bin/waitfor /go/bin/waitfor
CMD ["/go/bin/waitfor"]
