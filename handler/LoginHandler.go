package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"photo-sharing/db"
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

	fmt.Printf("Email: %s \n", email)
	fmt.Printf("Password: %s \n", password)

	db.DB.Find(&User{Email: email})

	return context.Render(http.StatusOK, "Login.html", echo.Map{})
}
