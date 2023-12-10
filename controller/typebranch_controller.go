// controller/TypeBranch_controller.go
package controller

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/labstack/echo/v4"
	"inventory_control/database"
	"inventory_control/model"
)

// GetTypeBranches godoc
//	@Summary		支部のリストを取得する
//	@Description	全ての支部情報を取得する
//	@Tags			types_branches
//	@Produce		json
//	@Success		200	{array}	model.TypeBranch
//	@Router			/types_branches [get]
func GetTypeBranchAll(c echo.Context) error {
	TypeBranch := []model.TypeBranch{}
	database.DB.Find(&TypeBranch)
	return c.JSON(http.StatusOK, TypeBranch)
}

// GetTypeBranch godoc
//	@Summary		特定のIDで支部を取得する
//	@Description	IDによる支部情報の取得
//	@Tags			types_branches
//	@Produce		json
//	@Param			id	path		int	true	"支部のID"
//	@Success		200	{object}	model.TypeBranch
//	@Router			/types_branches/{id} [get]
func GetTypeBranch(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	TypeBranch := model.TypeBranch{}
	database.DB.First(&TypeBranch, id)
	return c.JSON(http.StatusOK, TypeBranch)
}

// CreateTypeBranch godoc
//	@Summary		新しい支部を作成する
//	@Description	新しい支部情報の作成
//	@Tags			types_branches
//	@Accept			json
//	@Produce		json
//	@Param			TypeBranch	body		model.TypeBranch	true	"支部情報"
//	@Success		201			{object}	model.TypeBranch
//	@Router			/types_branches [post]
func CreateTypeBranch(c echo.Context) error {
	TypeBranch := model.TypeBranch{}
	c.Bind(&TypeBranch)
	database.DB.Save(&TypeBranch)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/types_branches/%d", TypeBranch.ID))
	return c.NoContent(http.StatusCreated)
}

// UpdateTypeBranch godoc
//	@Summary		特定のIDで支部を更新する
//	@Description	IDによる支部情報の更新
//	@Tags			types_branches
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int					true	"支部のID"
//	@Param			TypeBranch	body		model.TypeBranch	true	"支部情報"
//	@Success		200			{object}	model.TypeBranch
//	@Router			/types_branches/{id} [put]
func UpdateTypeBranch(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	TypeBranch := model.TypeBranch{}
	database.DB.First(&TypeBranch, id)
	c.Bind(&TypeBranch)
	database.DB.Save(&TypeBranch)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/types_branches/%d", TypeBranch.ID))
	return c.NoContent(http.StatusNoContent)
}

// DeleteTypeBranch godoc
//	@Summary		特定のIDで支部を削除する
//	@Description	IDによる支部情報の削除
//	@Tags			types_branches
//	@Produce		json
//	@Param			id	path		int		true	"支部のID"
//	@Success		204	{string}	string	"支部情報が削除されました"
//	@Router			/types_branches/{id} [delete]
func DeleteTypeBranch(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	TypeBranch := model.TypeBranch{}
	database.DB.Delete(&TypeBranch, id)
	return c.NoContent(http.StatusNoContent)
}
