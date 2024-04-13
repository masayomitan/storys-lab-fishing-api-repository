FROM mysql:8.0

# 環境変数を設定
ENV MYSQL_DATABASE=your_database \
    MYSQL_ROOT_PASSWORD=your_root_password \
    MYSQL_USER=your_user \
    MYSQL_PASSWORD=your_password

# カスタム設定ファイルをコンテナにコピー
# COPY custom.cnf /etc/mysql/conf.d/custom.cnf

# 初期化スクリプトをコンテナにコピー
# COPY init.sql /docker-entrypoint-initdb.d/

# ポート3306を公開
EXPOSE 3306

# ヘルスチェックを設定
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD mysqladmin ping -h localhost || exit 1

# デフォルトでmysqlサーバーを起動する
CMD ["mysqld"]
