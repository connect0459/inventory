package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"inventory_control/config"
	"inventory_control/controller"
	_ "inventory_control/docs/v1"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// v1のルーティング
	v1 := e.Group("/api/v1")
	{
		// 開発モード時のみSwaggerドキュメントを表示
		if config.AppConfig.DevMode {
			v1.GET("/swagger/*", echoSwagger.WrapHandler)
			// ドキュメントの生成は"swag init -o ./docs/v1"
			v1.File("/swagger/doc.json", "./docs/v1/swagger.json")
			v1.File("/swagger/doc.yaml", "./docs/v1/swagger.yaml")
		}

		// user
		v1.GET("/users", controller.GetUserAll)
		v1.GET("/users/:user_id", controller.GetUser)
		v1.POST("/users", controller.CreateUser)
		v1.PATCH("/users/:user_id", controller.UpdateUser)
		v1.DELETE("/users/:user_id", controller.DeleteUser)

		// books_info
		v1.GET("/books_info", controller.GetBookInfoAll)
		v1.GET("/books_info/:id", controller.GetBookInfo)
		v1.POST("/books_info", controller.CreateBookInfo)
		v1.PUT("/books_info/:id", controller.UpdateBookInfo)
		v1.DELETE("/books_info/:id", controller.DeleteBookInfo)

		// books_stocks
		v1.GET("/books_stocks", controller.GetBookStockAll)
		v1.GET("/books_stocks/:id", controller.GetBookStock)
		v1.POST("/books_stocks", controller.CreateBookStock)
		v1.PUT("/books_stocks/:id", controller.UpdateBookStock)
		v1.DELETE("/books_stocks/:id", controller.DeleteBookStock)

		// types_branches
		v1.GET("/types_branches", controller.GetTypeBranchAll)
		v1.GET("/types_branches/:id", controller.GetTypeBranch)
		v1.POST("/types_branches", controller.CreateTypeBranch)
		v1.PUT("/types_branches/:id", controller.UpdateTypeBranch)
		v1.DELETE("/types_branches/:id", controller.DeleteTypeBranch)

		// types_genres
		v1.GET("/types_genres", controller.GetTypeGenreAll)
		v1.GET("/types_genres/:id", controller.GetTypeGenre)
		v1.POST("/types_genres", controller.CreateTypeGenre)
		v1.PUT("/types_genres/:id", controller.UpdateTypeGenre)
		v1.DELETE("/types_genres/:id", controller.DeleteTypeGenre)
	}

	return e
}
