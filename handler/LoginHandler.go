package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
	"photo-sharing/util"
	"time"
)

type User struct {
	ID        uint
	Email     string
	Password  string
	CreatedAt time.Time
}

func GetLogin(context echo.Context) error {
	return context.Render(http.StatusOK, "Login.html", echo.Map{})
}

func PostLogin(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	user := &User{}
	db.DB.Where("email = ?", email).First(&user)
	if user == nil || !util.CheckPasswordHash(password, user.Password) {
		return context.Render(http.StatusUnauthorized, "Login.html", echo.Map{"error": "Invalid username & password combination"})
	}

	return context.Render(http.StatusOK, "Login.html", echo.Map{})
}
