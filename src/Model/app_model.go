package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"src/database"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jszwec/csvutil"
	"github.com/koron/go-dproxy"
	"github.com/olahol/melody"
	"golang.org/x/crypto/bcrypt"
)

func Userlogin(ctx *gin.Context) {
	reqUser := database.User{}
	//クライアントからのjsonデータをユーザー構造体にbinding
	err := ctx.BindJSON(&reqUser)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		log.Println(err)
	} else {

		pass := reqUser.Password
		user := database.GetUserData(reqUser)
		//DBから取得してきたpasswordはハッシュ値
		hashPass := user.Password

		//password比較
		err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))

		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			log.Println(err)
		} else {
			//sessionのセットアップ
			session := sessions.Default(ctx)

			//セッションにuserIDを格納
			log.Println("userdata:", user)
			log.Println("reqdata:", reqUser)
			sessionUser, err := json.Marshal(user)

			if err == nil {
				session.Set("loginUser", string(sessionUser))
				session.Save()

				log.Println("session Log", session.Get("loginUser"))
			} else {
				ctx.Status(http.StatusInternalServerError)
			}
		}
	}
}

func GetUserItem(ctx *gin.Context) {
	session := sessions.Default(ctx)

	//構造体インスタンスの生成
	//sessionから取ったユーザ情報の構造体へのマッピング
	user := sessionCheck(session, ctx)

	//dbからuserItem情報を取得してjson形式で返却
	userItems := database.GetUserItemData(user)
	log.Println(user, ":", userItems)

	ctx.JSON(200, userItems)
}

func PostUserItem(ctx *gin.Context) {
	session := sessions.Default(ctx)

	//構造体インスタンスの生成
	//sessionから取ったユーザ情報の構造体へのマッピング
	user := sessionCheck(session, ctx)

	var ItemDiff []database.UserItemJson

	err := ctx.BindJSON(&ItemDiff)

	log.Println(ItemDiff)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Abort()
	}

	database.SetUserItemData(user, ItemDiff)
}

func GetConsteData(ctx *gin.Context) {
	session := sessions.Default(ctx)

	//構造体インスタンスの生成
	//sessionから取ったユーザ情報の構造体へのマッピング
	user := sessionCheck(session, ctx)

	userconste := database.GetUserConstellationData(user)

	ctx.JSON(200, userconste)
}

func PostConsteData(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := sessionCheck(session, ctx)

	var ConsteData database.UserConstellationJson

	//なぜかbindJsonで出来なかったので直接ボディを読んでバインディングを実行
	/*buf := make([]byte, 2048)
	n, _ := ctx.Request.Body.Read(buf)
	b := string(buf[0:n])
	log.Println("string:", b)
	b = fmt.Sprintf("{%s}", b)
	err := json.Unmarshal([]byte(b), &ConsteData)*/

	//binding
	err := ctx.BindJSON(&ConsteData)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		log.Println("err:", err)
		ctx.Abort()
	}

	log.Println("ModelConste", ConsteData)
	err = database.CreateUserConstellationData(user, ConsteData)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		log.Println("primary key:", err)
		ctx.String(400, "もう一度試してください") //クライアントにここは実装したいですと伝える
	}

}

func QuizeGet(ctx *gin.Context) {
	var Quize []database.QuizeDataJson

	//csvファイル読み込み
	b, err := ioutil.ReadFile("database/quize.csv")

	if err != nil {
		log.Println("csv read err:", err)
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}

	//csvデータを構造体にマッピング
	err = csvutil.Unmarshal(b, &Quize)

	if err != nil {
		log.Println("Unmarshal err:", err)
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}

	//乱数の生成
	rand.Seed(time.Now().UnixNano())

	i := rand.Intn(100) % len(Quize)

	ctx.JSON(200, Quize[i])

}

func sessionCheck(session sessions.Session, ctx *gin.Context) database.User {
	user := database.User{}

	userJson, err := dproxy.New(session.Get("loginUser")).String()
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
	}
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
	}
	return user
}

func EventFlag(ctx *gin.Context) {
	m := melody.New()
	//クライアントIpが127.0.0.1ならばブロードキャスト実行
	cliantIp := ctx.ClientIP()

	if cliantIp == "127.0.0.1" {
		//イベントフラグをws接続中の全クライアントに送信、値はuint8の「1」
		m.Broadcast([]byte{1})
	} else {
		ctx.Status(http.StatusUnauthorized)
		ctx.String(401, "ここにはlocalhost以外アクセス出来ません")
	}
}

func WsConnect(ctx *gin.Context) {
	//ws接続確立
	m := melody.New()

	err := m.HandleRequest(ctx.Writer, ctx.Request)

	if err != nil {
		log.Println(err)
	}
}
