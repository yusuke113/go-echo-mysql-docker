# golangの最新
FROM golang:1.20.7

RUN apt-get update && \
  apt-get install -y git

WORKDIR /go/src

# ホストのgo.modとgo.sumをコンテナの作業ディレクトリにコピー
COPY /app/go.mod .
COPY /app/go.sum .

RUN go mod tidy
