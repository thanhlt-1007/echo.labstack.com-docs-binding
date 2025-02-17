package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID int `query:"id"`
}

func getUserHandler(context echo.Context) error {
	var user User
	err := context.Bind(&user)
	if err != nil {
		return context.JSON(
			http.StatusBadRequest,
			err,
		)
	}

	return context.JSON(
		http.StatusOK,
		user,
	)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/user", getUserHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
