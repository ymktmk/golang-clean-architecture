# Golang Clean Arc

自前で認証・認可サーバーを作成する
パスワードをハッシュ化する

## マイグレーション
```
docker exec -it todo go run migrate/migrate.go
```

## サーバー起動
```
docker-compose exec app go run main.go
```


## Create
```
curl -X POST http://localhost:9000/users/create \
-H 'Content-Type: application/json' \
-d '{"name": "tomoki", "email": "tt@gmail.com"}'
```