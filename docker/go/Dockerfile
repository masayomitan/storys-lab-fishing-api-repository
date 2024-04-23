# ステージ1: ビルド環境
FROM golang:1.22.1-alpine AS builder

WORKDIR /

# Airとビルドツールのインストール
RUN apk add --no-cache go ca-certificates
RUN go install github.com/cosmtrek/air@v1.51.0
RUN apk update &&  apk add git

# 依存関係のインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./

# ステージ2: 実行環境
FROM alpine:latest
RUN apk --no-cache add ca-certificates

EXPOSE 80

# ビルドした実行ファイルのコピー
COPY --from=builder /main .
COPY --from=builder /.env .
COPY --from=builder /go/bin/air /usr/local/bin/

# Airの設定ファイルのコピー
COPY .air.toml .air.toml

# 環境に応じた実行コマンドの設定
CMD ["air", "-c", ".air.toml"]
# CMD if [ "$ENVIRONMENT" = "development" ]; then air -c .air.toml; else ./main; fi
# CMD ["./main"]
