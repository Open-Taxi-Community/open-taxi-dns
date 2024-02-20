FROM golang:1.18 as build

WORKDIR /go/src/dnsapi

COPY ../../cmd/dnsapi .
COPY cmd/dnsapi .
ADD ../../cmd/dnsapi  /go/src/github.com/Open-Taxi-Community/open-taxi-dns
RUN GOFLAGS=-mod=mod go generate ./...

#RUN go build -v -o dnsapi .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64\
    go build -a -installsuffix cgo \
    -o /go/bin/dnsapi

RUN chmod +x /go/bin/dnsapi

FROM alpine:3
RUN apk add --no-cache ca-certificates
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip

COPY --from=build /go/bin/dnsapi/. /dnsapi

ENV PORT 8080

ENTRYPOINT ["/dnsapi", "run"]