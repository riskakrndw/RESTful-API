package main

import (
	"projects/config"
	"projects/controllers"
	"projects/models"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	db, err := gorm.Open(mysql.Open(config.DB_CONNECTION_STRING))
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	userModel := models.NewGormUserModel(db)
	userController := controllers.CreateGetUserController(userModel)
	e.GET("/user", userController)
	e.Logger.Fatal(e.Start(":8000"))
}
