# storys-lab-fishing-api-repository

# Docker

### ビルド
``` docker command
docker-compose up -d --build
```

### ローカルでAir使用
```
docker-compose up
```

# API仕様書
## swaggo
https://github.com/swaggo/swag



# マイグレーション
## golang-migrate 
https://github.com/golang-migrate/migrate


### インストール
```
brew install golang-migrate
```

### マイグレーションの作成
```
migrate create -ext sql -dir app/infrastructure/database/migrations -seq create_table_name
```

### マイグレーションの実行
```
migrate -path app/infrastructure/database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" up
```

### マイグレーションのロールバック
```
migrate -path app/infrastructure/database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" down 1
```

### 全てのマイグレーションをロールバック
```
migrate -path app/infrastructure/database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" down -all
```

### 特定のバージョンへのマイグレーション
```
migrate -path app/infrastructure/database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" goto 2
```

### 現在のバージョンの確認
```
migrate -path app/infrastructure/database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" version
```
