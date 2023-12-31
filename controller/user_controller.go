package controller

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	. "inventory_control/model"
	. "inventory_control/database"
)

func GetUserAll(c echo.Context) error {
	user := []User{}

	if err := DB.Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "users not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {
	user := User{}

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "user not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	user := new(User)

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := DB.Model(&user).Updates(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func DeleteUser(c echo.Context) error {
	user := new(User)

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.Where("id = ?", userID).Delete(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, user)
}
