FROM golang:1.18 as build

WORKDIR /go/src/server

COPY ../../internal/dnsapi  .
COPY internal/dnsapi/server .
ADD ../../internal/dnsapi  /go/src/github.com/Open-Taxi-Community/open-taxi-dns
RUN GOFLAGS=-mod=mod go generate ./...

#RUN go build -v -o server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64\
    go build -a -installsuffix cgo \
    -o /go/bin/server

RUN chmod +x /go/bin/server

FROM alpine:3
RUN apk add --no-cache ca-certificates
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip

COPY --from=build /go/bin/server/. /server

ENV PORT 8080

ENTRYPOINT ["/server", "run"]