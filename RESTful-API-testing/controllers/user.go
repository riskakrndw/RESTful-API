package controllers

import (
	"projects/models"

	"github.com/labstack/echo"
)

func CreateGetUserController(userModel models.UserModel) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := userModel.Get()
		return c.JSON(200, users)
	}
}
