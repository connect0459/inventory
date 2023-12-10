# 開発環境用
FROM golang:1.21.4-bullseye

# 環境変数
ENV ROOT=/go/src/app
WORKDIR ${ROOT}

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . .

# アップデートとgitのインストール
RUN apt-get update && \
    apt-get install -y git && \
    apt-get -y install default-mysql-client

# go.modをもとにして依存関係をダウンロード
RUN go mod download && go mod verify && go mod tidy
