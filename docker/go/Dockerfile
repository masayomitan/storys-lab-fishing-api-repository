# ステージ1: ビルド環境
FROM golang:1.22.1-alpine AS builder

# Airとビルドツールのインストール
RUN apk add --no-cache ca-certificates && \
    wget -O - -q https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /bin

RUN apk update &&  apk add git
WORKDIR /app

# 依存関係のインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./app

# ステージ2: 実行環境
FROM alpine:latest  
RUN apk --no-cache add ca-certificates


# ビルドした実行ファイルのコピー
COPY --from=builder /app/main .

# Airの設定ファイルのコピー
COPY .air.toml .air.toml

# 実行
CMD ["air", "-c", ".air.toml"]
# CMD ["./main"]
