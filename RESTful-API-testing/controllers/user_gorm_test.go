package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projects/config"
	"projects/lib/database"
	"projects/models"
	"testing"

	"github.com/labstack/echo"
)

func TestDBGetUserController(t *testing.T) {
	//bikin db
	db, err := database.CreateDB(config.TEST_DB_CONNECTION_STRING)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&models.User{})
	db.Delete(&models.User{}, "1=1")
	m := models.NewGormUserModel(db)

	//bikin controller
	userController := CreateGetUserController(m)
	if err != nil {
		t.Error(err)
	}

	//insert new data
	m.Insert(models.User{Name: "Riska", Email: "tes@gmail.com", Password: "123"})

	//test controller
	testGetUserController(t, userController)
	db.Delete(&models.User{}, "1=1")
}

func testGetUserController(t *testing.T, userController echo.HandlerFunc) {
	//coba request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	userController(c)

	//test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}

	body := rec.Body.Bytes()
	var users []models.User
	if err := json.Unmarshal(body, &users); err != nil {
		t.Error(err)
	}
	if len(users) != 1 {
		t.Errorf("expected one user, got: %#v", users)
		return
	}
	if users[0].Name != "Riska" {
		t.Errorf("expected Riska, got: %#v", users[0].Name)
	}
}
