# backend
わかんないッピのバックエンド用リポジトリ

Model

DB: DBとやり取りするファイルを置く場所

API: フロントとDBをつなぐ場所

auth: ユーザー認証する場所

pub: イベント処理

Controller

rout: ルーティングする場所(アプリ側にポートの案内をする)

Viewはなし

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