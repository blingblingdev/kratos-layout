FROM golang:1.15 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /data/conf

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
#VOLUME /data/conf

CMD ["./github.com/ffy/kratos-layout", "-conf", "/data/conf"]
