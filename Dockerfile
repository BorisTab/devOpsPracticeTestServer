FROM --platform=linux/amd64 golang:1.19.1-alpine3.16 as builder

COPY . /build

WORKDIR /build

RUN go build -o /tmp/back_server ./main.go

FROM alpine:3.16

COPY --from=builder /tmp/back_server /usr/bin/back_server

RUN chmod +x /usr/bin/back_server

ENTRYPOINT ["back_server"]