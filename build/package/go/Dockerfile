# Goバージョン
FROM golang:1.21.4-bullseye

# 環境変数
# ENV ROOT=/go/src/inventory_control
ENV ROOT=/go/src/app
# ENV CGO_ENABLED 0
WORKDIR ${ROOT}

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . .

# アップデートとgitのインストール
RUN apt-get update && \
    apt-get install -y git && \
    apt-get -y install default-mysql-client

# Userの作成
# ENV USERNAME = developer
# RUN useradd -m -d /home/${USERNAME} -s /bin/bash ${USERNAME}
# USER ${USERNAME}
# RUN git config --global user.email "you@example.com"
# RUN git config --global user.name "Your Name"

# Goモジュールを有効化
# goのプロジェクトがGOPATHの配下になければ=on
# RUN go env -w GO111MODULE=on

# コマンドのインストール
# RUN go install github.com/cweill/gotests/gotests@v1.6.0 && \
#     go install github.com/fatih/gomodifytags@v1.16.0 && \
#     go install github.com/josharian/impl@v1.1.0 && \
#     go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 && \
#     go install github.com/go-delve/delve/cmd/dlv@latest && \
#     go install honnef.co/go/tools/cmd/staticcheck@latest && \
#     go install golang.org/x/tools/gopls@latest

# パッケージのインポートと依存関係の取得
# RUN go get github.com/labstack/echo/v4 && \
#     go get github.com/labstack/echo/v4/middleware && \
#     go get github.com/swaggo/echo-swagger && \
#     go get gorm.io/gorm && \
#     go get gorm.io/driver/mysql

RUN go mod download && go mod verify

# # Goアプリケーションのビルド
# RUN go build -o app .

# # サーバーをポート7000で実行
# EXPOSE 7000

# # コンテナを起動時に実行するコマンド
# CMD ["./app"]
