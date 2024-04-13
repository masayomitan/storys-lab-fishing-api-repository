# storys-lab-fishing-api-repository

## 

``` dockert command
docker-compose up -d --build
```

# golang-migrate

## goマイグレーション
```
brew install golang-migrate
```

## マイグレーションの作成
```
migrate create -ext sql -dir database/migrations -seq create_table_name
```

## マイグレーションの実行
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" up
```

## マイグレーションのロールバック
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" down 1
```

## 全てのマイグレーションをロールバック
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" down -all
```

## 特定のバージョンへのマイグレーション
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" goto 2
```

## 現在のバージョンの確認
```
migrate -path database/migrations -database "mysql://root@tcp(127.0.0.1:3306)/story_fishing_db" version
```
