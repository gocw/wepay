package main

import (
	"net/http"

	"github.com/labstack/echo"
	"errors"
)

func initRouter() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "wepay")
	})
	e.POST("/wepay/api",api)
	e.Logger.Fatal(e.Start(":1323"))
}

