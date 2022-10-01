# backend
わかんないッピのバックエンド用リポジトリ



# ルーティングについて
routingはlogin済みのユーザーのみアクセスできるAPIがあります
```go:routing.go

    //だれでもアクセスできるAPI
    r.POST("/login", model.Userlogin)   

	AuthUserGroup := r.Group("/auth")
	AuthUserGroup.Use(middleware.LoginCheck)
	{
        //ログイン済みでないとアクセスできないAPI群
		AuthUserGroup.GET("/OrigConste/Get", model.GetConsteData)
		AuthUserGroup.GET("/UserItem/Get", model.GetUserItem)      
		AuthUserGroup.POST("/UserItem/Post", model.PostUserItem)  
		AuthUserGroup.GET("/Quize/Get", model.QuizeGet)            
		AuthUserGroup.POST("/OrigConste/Post", model.PostConsteData) 
	}

```

# 各種API
## loginAPI
### 概要
ユーザーloginする時のAPIです<br>
sessionでログインユーザーを管理するので、クライアントはcookie管理ができるライブラリなりを使ってください<br>
### アクセスURL
```URL
http://serverName:port/login
```
### リクエストbody(Json)
```
{"userId":"User", "userPass":"Password"}
```
### httpステータス<br>
httpレスポンスステータスコードは以下の3種類です<br>
クライアントのエラーハンドリングに役立ててください<br>
| httpstatus | 状態 |
| ---------- | ---  |
|    200     | 成功 |
|    400     | login失敗<br>(passwordが違う,等)|
|    500     | login失敗<br>(サーバ側の問題)|
|    404     | URLが見つからない<br>|
<br><br>

## オリジナル星座取得API
### 概要
ユーザーのオリジナル星座を一括で取得してJSON形式でリクエストクライアントに返すAPIです<br>
sessionでログインユーザーを管理するので、クライアントはcookie管理ができるライブラリなりを使ってください<br>

### アクセスURL
```URL
http://serverName:port/auth/OrigConste/Get
```
### レスポンスbody(Json)
```
[
    {"consteId":"conste1", "consteName":"aiueo", "consteData"://保存データが決まってません.決まったら書き加えます},

    {"consteId":"conste2", "consteName":"aiueo", "consteData"://保存データが決まってません.決まったら書き加えます}
]
```

### httpステータス<br>
httpレスポンスステータスコードは以下の3種類です<br>
クライアントのエラーハンドリングに役立ててください<br>
| httpstatus | 状態 |
| ---------- | ---  |
|    200     | 成功 |
|    401     | 認証情報が切れている<br>|
|    500     | 失敗<br>(サーバ側の問題)|
<br><br><br>




# docker環境使い方
## 設定ファイル
backend/docker/mysql/db.envを作成し、下のテンプレを保存
```
MYSQL_DATABASE=wakannaippi
MYSQL_ROOT_PASSWORD=****(任意の文字列)
MYSQL_HOST=wakannaippi_db
TZ='Asia/Tokyo'
```
## 起動
backend/dockerに移動して、下のコマンドを実行
```
docker-compose up -d
```

## サーバアプリケーション実行
コンテナ内のshellに移動
```
docker exec -it wakannaippi_go sh
```
ユーザーセットアップ<br>
src/UserSetup  [コンテナ内]
```
go run main.go
```
サーバ起動<br>
src/cmd [コンテナ内]
```
go run main.go
```
cronを使う<br>
src/cron [コンテナ内]
```
go run main.go
```
