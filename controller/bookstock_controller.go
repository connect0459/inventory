package controller

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/labstack/echo/v4"
	"inventory_control/model"
	"inventory_control/database"
)

// GetBookStocks godoc
//	@Summary		書籍在庫のリストを取得する
//	@Description	全ての書籍在庫情報を取得する
//	@Tags			books_stocks
//	@Produce		json
//	@Success		200	{array}	model.BookStock
//	@Router			/books_stocks [get]
func GetBookStockAll(c echo.Context) error {
	bookStock := []model.BookStock{}
	database.DB.Find(&bookStock)
	return c.JSON(http.StatusOK, bookStock)
}

// GetBookStock godoc
//	@Summary		特定のIDで書籍在庫を取得する
//	@Description	IDによる書籍在庫情報の取得
//	@Tags			books_stocks
//	@Produce		json
//	@Param			id	path		int	true	"書籍在庫のID"
//	@Success		200	{object}	model.BookStock
//	@Router			/books_stocks/{id} [get]
func GetBookStock(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookStock := model.BookStock{}
	database.DB.First(&bookStock, id)
	return c.JSON(http.StatusOK, bookStock)
}

// CreateBookStock godoc
//	@Summary		新しい書籍在庫を作成する
//	@Description	新しい書籍在庫情報の作成
//	@Tags			books_stocks
//	@Accept			json
//	@Produce		json
//	@Param			BookStock	body		model.BookStock	true	"書籍在庫情報"
//	@Success		201			{object}	model.BookStock
//	@Router			/books_stocks [post]
func CreateBookStock(c echo.Context) error {
	bookStock := model.BookStock{}
	c.Bind(&bookStock)
	database.DB.Save(&bookStock)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books_stocks/%d", bookStock.ID))
	return c.NoContent(http.StatusCreated)
}

// UpdateBookStock godoc
//	@Summary		特定のIDで書籍在庫を更新する
//	@Description	IDによる書籍在庫情報の更新
//	@Tags			books_stocks
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int				true	"書籍在庫のID"
//	@Param			BookStock	body		model.BookStock	true	"書籍在庫情報"
//	@Success		200			{object}	model.BookStock
//	@Router			/books_stocks/{id} [put]
func UpdateBookStock(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookStock := model.BookStock{}
	database.DB.First(&bookStock, id)
	c.Bind(&bookStock)
	database.DB.Save(&bookStock)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books_stocks/%d", bookStock.ID))
	return c.NoContent(http.StatusNoContent)
}

// DeleteBookStock godoc
//	@Summary		特定のIDで書籍在庫を削除する
//	@Description	IDによる書籍在庫情報の削除
//	@Tags			books_stocks
//	@Produce		json
//	@Param			id	path		int		true	"書籍在庫のID"
//	@Success		204	{string}	string	"書籍在庫情報が削除されました"
//	@Router			/books_stocks/{id} [delete]
func DeleteBookStock(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookStock := model.BookStock{}
	database.DB.Delete(&bookStock, id)
	return c.NoContent(http.StatusNoContent)
}
