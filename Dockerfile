## Compile stage
#FROM golang:1.20 AS build-env
#
#ADD . /dockerdev
#WORKDIR /dockerdev
#
#RUN go build -o /server
#
## Final stage
#FROM alpine:latest
#
#EXPOSE 8000
#
#WORKDIR /
#COPY --from=build-env /server /
#COPY --from=build-env /config.yaml /
#
#CMD ["/server"]

FROM golang:1.20 AS builder

COPY . /src
WORKDIR /src

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

RUN go build -o main main.go

FROM ubuntu

RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN  apt-get clean
RUN apt-get update
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase

VOLUME /home

COPY --from=builder /src/main /home
COPY --from=builder /src/config.yaml /home

WORKDIR /home

CMD ["./main"]
