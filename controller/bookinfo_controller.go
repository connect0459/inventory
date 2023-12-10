// controller/bookinfo_controller.go
package controller

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/labstack/echo/v4"
	"inventory_control/database"
	"inventory_control/model"
)

//	@Summary		全ての書籍情報を取得
//	@Description	全ての書籍情報を取得する
//	@Tags			books_info
//	@Produce		json
//	@Success		200	{array}	model.BookInfo
//	@Router			/books [get]
func GetBookInfoAll(c echo.Context) error {
	bookInfo := []model.BookInfo{}
	database.DB.Find(&bookInfo)
	return c.JSON(http.StatusOK, bookInfo)
}

//	@Summary		書籍情報を取得
//	@Description	IDによる書籍情報の取得
//	@Tags			books_info
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	model.BookInfo
//	@Router			/books/{id} [get]
func GetBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.First(&bookInfo, id)
	return c.JSON(http.StatusOK, bookInfo)
}

//	@Summary		書籍情報を作成
//	@Description	新しい書籍情報の作成
//	@Tags			books_info
//	@Accept			json
//	@Produce		json
//	@Param			bookInfo	body		model.BookInfo	true	"Book Info"
//	@Success		200			{object}	model.BookInfo
//	@Router			/books [post]
func CreateBookInfo(c echo.Context) error {
	bookInfo := model.BookInfo{}
	c.Bind(&bookInfo)
	database.DB.Save(&bookInfo)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books_info/%d", bookInfo.ID))
	return c.NoContent(http.StatusCreated)
}

//	@Summary		書籍情報を更新
//	@Description	IDによる書籍情報の更新
//	@Tags			books_info
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int				true	"Book ID"
//	@Param			bookInfo	body		model.BookInfo	true	"Book Info"
//	@Success		200			{object}	model.BookInfo
//	@Router			/books/{id} [put]
func UpdateBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.First(&bookInfo, id)
	c.Bind(&bookInfo)
	database.DB.Save(&bookInfo)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books_info/%d", bookInfo.ID))
	return c.NoContent(http.StatusNoContent)
}

//	@Summary		書籍情報を削除
//	@Description	IDによる書籍情報の削除
//	@Tags			books_info
//	@Produce		json
//	@Param			id	path		int		true	"Book ID"
//	@Success		200	{string}	string	"BookInfo is deleted"
//	@Router			/books/{id} [delete]
func DeleteBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.Delete(&bookInfo, id)
	return c.NoContent(http.StatusNoContent)
}
