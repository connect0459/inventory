package main

import (
	// "log"
	// "github.com/joho/godotenv"
	"inventory_control/database"
	"inventory_control/route"
)

//	@title			inventory_control API
//	@version		1.0
//	@description	在庫管理を行うAPIサーバー
//	@contact.name	API Support
//	@contact.email	connect0459@github.com
//	@host			localhost:7000
//	@BasePath		/api/v1
func main() {
	// データベース接続を取得
	db, _ := database.DB.DB()
	// データベース接続を閉じる
	defer db.Close()

	// ルーターを作成
	e := route.Router()
	// サーバーを起動し、エラーが発生した場合はログに記録
	e.Logger.Fatal(e.Start(":7000"))
}