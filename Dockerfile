FROM --platform=linux/amd64 golang:1.19.1-alpine3.16 as builder

COPY . /build

WORKDIR /build

RUN go build -o /tmp/server .

FROM alpine:3.16

COPY --from=builder /tmp/server /usr/bin/server

RUN chmod +x /usr/bin/server

ENTRYPOINT ["server"]