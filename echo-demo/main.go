package main

/*
   @Auth: menah3m
   @Desc:
*/

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	//e.HideBanner = true

	e.Logger.Fatal(e.Start(":8888"))
}
