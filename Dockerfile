FROM golang:1.12.7 as build

WORKDIR /go/src/github.com/abramd/strange-server
COPY ./ ./

RUN GO111MODULE=on go mod vendor

RUN go build -a -o /go/bin/app ./cmd/strange

#FROM scratch
FROM debian:stretch-slim

COPY --from=build /go/bin/app /app
COPY --from=build /go/src/github.com/abramd/strange-server/internal/db/migrations /migrations
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /
ENTRYPOINT ["/app"]