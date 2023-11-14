package main

import (
	// "github.com/sirupsen/logrus"
	// "github.com/valyala/fasthttp"
	"inventory_control/routes"
	"inventory_control/database"
)

func main() {
	db, _ := database.DB.DB()
	defer db.Close()

	e := apiRoute.NewRouter()
	e.Logger.Fatal(e.Start(":7000"))
}
