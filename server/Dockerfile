FROM golang:1.17.8-alpine3.15

RUN mkdir /go/src/server && apk update && apk add git

WORKDIR /go/src/server

# ホストのファイル・ディレクトリをコンテナイメージにコピー
ADD . /go/src/server

RUN go get github.com/gin-gonic/gin

# RUN go mod download

# ホットリロード用にair
# RUN go build -o /go/bin/air github.com/cosmtrek/air

# CMD ["air", "-c", ".air.toml"]