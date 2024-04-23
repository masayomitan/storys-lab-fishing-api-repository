# storys-lab-fishing-api-repository

# Docker

### ビルド
``` dockert command
docker-compose up -d --build
```

### ローカルでAir使用
```
docker-compose up
```

# API仕様書
## swaggo https://github.com/golang-migrate/migrate](https://github.com/swaggo/swag



# マイグレーション
## golang-migrate https://github.com/golang-migrate/migrate


### インストール
```
brew install golang-migrate
```

### マイグレーションの作成
```
migrate create -ext sql -dir database/migrations -seq create_table_name
```

### マイグレーションの実行
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" up
```

### マイグレーションのロールバック
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" down 1
```

### 全てのマイグレーションをロールバック
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" down -all
```

### 特定のバージョンへのマイグレーション
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" goto 2
```

### 現在のバージョンの確認
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" version
```
