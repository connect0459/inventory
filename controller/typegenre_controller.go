// controller/TypeGenre_controller.go
package controller

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/labstack/echo/v4"
	"inventory_control/database"
	"inventory_control/model"
)

// GetTypeGenres godoc
//	@Summary		ジャンルのリストを取得する
//	@Description	全てのジャンル情報を取得する
//	@Tags			types_genres
//	@Produce		json
//	@Success		200	{array}	model.TypeGenre
//	@Router			/types_genres [get]
func GetTypeGenreAll(c echo.Context) error {
	TypeGenre := []model.TypeGenre{}
	database.DB.Find(&TypeGenre)
	return c.JSON(http.StatusOK, TypeGenre)
}

// GetTypeGenre godoc
//	@Summary		特定のIDでジャンルを取得する
//	@Description	IDによるジャンル情報の取得
//	@Tags			types_genres
//	@Produce		json
//	@Param			id	path		int	true	"ジャンルのID"
//	@Success		200	{object}	model.TypeGenre
//	@Router			/types_genres/{id} [get]
func GetTypeGenre(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	TypeGenre := model.TypeGenre{}
	database.DB.First(&TypeGenre, id)
	return c.JSON(http.StatusOK, TypeGenre)
}

// CreateTypeGenre godoc
//	@Summary		新しいジャンルを作成する
//	@Description	新しいジャンル情報の作成
//	@Tags			types_genres
//	@Accept			json
//	@Produce		json
//	@Param			TypeGenre	body		model.TypeGenre	true	"ジャンル情報"
//	@Success		201			{object}	model.TypeGenre
//	@Router			/types_genres [post]
func CreateTypeGenre(c echo.Context) error {
	TypeGenre := model.TypeGenre{}
	c.Bind(&TypeGenre)
	database.DB.Save(&TypeGenre)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/types_genres/%d", TypeGenre.ID))
	return c.NoContent(http.StatusCreated)
}

// UpdateTypeGenre godoc
//	@Summary		特定のIDでジャンルを更新する
//	@Description	IDによるジャンル情報の更新
//	@Tags			types_genres
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int				true	"ジャンルのID"
//	@Param			TypeGenre	body		model.TypeGenre	true	"ジャンル情報"
//	@Success		204			{object}	model.TypeGenre
//	@Router			/types_genres/{id} [put]
func UpdateTypeGenre(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	TypeGenre := model.TypeGenre{}
	database.DB.First(&TypeGenre, id)
	c.Bind(&TypeGenre)
	database.DB.Save(&TypeGenre)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/types_genres/%d", TypeGenre.ID))
	return c.NoContent(http.StatusNoContent)
}

// DeleteTypeGenre godoc
//	@Summary		特定のIDでジャンルを削除する
//	@Description	IDによるジャンル情報の削除
//	@Tags			types_genres
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"ジャンルのID"
//	@Success		204	{string}	string	"ジャンル情報が削除されました"
//	@Router			/types_genres/{id} [delete]
func DeleteTypeGenre(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	TypeGenre := model.TypeGenre{}
	database.DB.Delete(&TypeGenre, id)
	return c.NoContent(http.StatusNoContent)
}
