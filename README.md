# Goを使ってUserのCURD処理 APIサーバー

# Entities

これを中心にして開発するのがCleanArchitecture
Userというモデルがどんなデータを持ち、どんな振る舞いを持つのか。
このDomainを中心にしてアーキテクチャを考える。

```
type User struct {
    ID        int    `json:"id"`
    Name string `json:"name"`
}

type Users []User

```


# Infrastructure層

DBとのアクセスやHTTPリクエストを受け取る
次の層にSqlhandlerも渡してあげる

# Interface層

### Gateways
内部 ←→ 外部(DB、FileStorage 等の Device)
### Controller
外部(クライアント) → 内部
ここからusecaseに渡す
### Presenters
内部 → 外部(クライアント)
usecaseから受け取りresponseを渡す

Infrastructure層からのリクエストをここで加工して
ビジネスロジックであるUseCace層に渡す。
https://tech-blog.optim.co.jp/entry/2019/01/29/173000

# UseCases層







# example

Infrastructure層でルーティング書く(HTTP Handler的な役割)
 → Interface層のController リクエストボディを取得 Entity 構造体に変換
 → Usecase層のRepositoryでSQL実行
 → Interface層のPresenters  Jsonに変換してレスポンスを行う
 