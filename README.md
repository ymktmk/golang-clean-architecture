# Golangクリーンアーキテクチャ

## Create
```
curl -X POST http://localhost:8080/users/create -H 'Content-Type: application/json' -d '{"name":"yyy"}'
```

## Show
```
curl http://localhost:8080/users/1
```

## Index
```
curl http://localhost:8080/users
```

jwt認証付き
バリデーション