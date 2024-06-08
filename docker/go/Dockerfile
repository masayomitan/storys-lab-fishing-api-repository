# ステージ1: ビルド環境
FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Airとビルドツールのインストール
RUN apk add --no-cache go ca-certificates
RUN go install github.com/cosmtrek/air@v1.51.0
RUN apk update &&  apk add git

# 依存関係のインストール
COPY go.mod ./
COPY go.sum ./

# モジュールのダウンロード
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./

# ステージ2: 実行環境
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Goをインストール
RUN apk add --no-cache go

# Goのパスを環境変数に設定
ENV PATH="/usr/local/go/bin:${PATH}"

# ビルドした実行ファイルのコピー
COPY --from=builder /app/main ./
COPY --from=builder /app/.env ./
COPY --from=builder /app/public ./public
COPY --from=builder /go/bin/air /usr/local/bin/
COPY --from=builder /app/app/infrastructure/database/migrations /app/infrastructure/database/migrations

# Airの設定ファイルのコピー
COPY .air.toml .air.toml

# 環境に応じた実行コマンドの設定
# CMD ["air"]
CMD ["./main"]
