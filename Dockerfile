FROM golang:1.11.1-alpine3.7

ENV GOPATH=/usr/share/golang
RUN apk update && apk add git bash && \
    git clone https://github.com/go-zkcli/zkcli /home && \
    go get github.com/codegangsta/cli && \
    go get github.com/go-zkcli/zkcli/output && \
    go get github.com/go-zkcli/zkcli/zk && \
    go get github.com/outbrain/golib/log && \
    cd /home/ && sh build.sh && \ 
    ls /tmp/zkcli/ && \ 
    mv /tmp/zkcli/zkcli /usr/bin/zkcli && \
    chmod +x /usr/bin/zkcli 