FROM golang:1.16.2 AS builder

RUN mkdir /vmware

COPY .. /vmware

RUN go env

WORKDIR /vmware

RUN go mod tidy -v

RUN go build -o metrics_collector ./cmd

FROM alpine:latest

RUN mkdir vmware

WORKDIR vmware

COPY --from=builder  /vmware/metrics_collector .

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD ["./metrics_collector"]
