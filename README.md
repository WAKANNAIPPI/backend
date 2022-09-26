# backend
わかんないッピのバックエンド用リポジトリ

Model

DB: DBとやり取りするファイルを置く場所

API: フロントとDBをつなぐ場所

auth: ユーザー認証する場所

pub: イベント処理

Controller

rout: ルーティングする場所(modelのルーティングをする)

Viewはなし<br>

# dbについて
database/db.iniを作成します<br>
--------------------------------<br>
[db]<br>
dbDriver = mysql <br>
dbName = wakannaippi <br>
dbUserName = ** <br>
dbUserPass = ** <br>
dbHost = localhost <br>
dbPort = 3306 <br>
--------------------------------<br>
のように書いてください

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
ユーザーloginする時のAPIです<br><br>
アクセスURL
```URL
http://serverName:8080/login
```
リクエストbody(Json)
```
{"userId":"User", "userPass":"Password"}
```
httpステータス<br>
リクエスト時、httpレスポンスステータスコードは以下の3種類です<br>
クライアントのエラーハンドリングに役立ててください<br>
| httpstatus | 状態 |
| ---------- | ---  |
|    200     | 成功 |
|    400     | login失敗<br>(passwordが違う,等)|
|    500     | login失敗<br>(サーバ側の問題)|



