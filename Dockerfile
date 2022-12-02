FROM golang:1.19.1-alpine3.16 as builder

WORKDIR /home
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o /home/example_rpc example.go

FROM alpine:latest

WORKDIR /home

COPY --from=builder /home/example_rpc ./
COPY --from=builder /home/etc/example.yaml ./

EXPOSE 9101
ENTRYPOINT ./example_rpc -f example.yaml