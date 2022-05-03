# Golang Clean Architecture

https://qiita.com/x-color/items/24ff2491751f55e866cf

## 認証

ログイン時にJWT Tokenを生成し、Cookieに付与してレスポンスを返す
TokenからUserIDを取得できるので、そのIDをキーにMySQLを検索して、User情報を取得する

## envファイル生成 & 記述

```
cp .env.example .env
```

## マイグレーション
```
docker exec -it todo go run migrate/migrate.go
```

## サーバー起動
```
docker-compose exec app go run main.go
```

# エンドポイント

## Create User
```
curl -X POST http://localhost:9000/users/create \
-H 'Content-Type: application/json' \
-d '{"name": "tomoki", "email": "tt@gmail.com", "password": "Tomomki0901"}'
```

## Get User

```
curl http://localhost:9000/user
```

## テスト

sqlmockを使っており「こういうSQLが来た時はこういう結果を返す」という指定ができる。
テスト用のDBを用意しなくても良いというメリットがある。