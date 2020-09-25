# Build Stage
FROM golang:1.14 AS build-stage

LABEL APP="build-gbox"
LABEL REPO="flamefatex/gbox"

ADD . /go/src/github.com/flamefatex/gbox
WORKDIR /go/src/github.com/flamefatex/gbox

RUN make build-alpine

# Final Stage
FROM alpine:3.12

ARG GIT_COMMIT
ARG VERSION
ARG APP_NAME

LABEL REPO="flamefatex/gbox"
LABEL GIT_COMMIT=${GIT_COMMIT}
LABEL VERSION=${VERSION}
LABEL APP_NAME=${APP_NAME}

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --no-cache tcpdump lsof net-tools tzdata curl dumb-init libc6-compat
RUN echo "hosts: files dns" > /etc/nsswitch.conf

ENV TZ Asia/Shanghai
ENV PATH=$PATH:/opt/gbox/bin

WORKDIR /opt/gbox/bin
# 配置文件

COPY --from=build-stage /go/src/github.com/flamefatex/gbox/bin/gbox .
RUN chmod +x /opt/gbox/bin/gbox

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/gbox/bin/gbox"]