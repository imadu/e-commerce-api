package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	mux := echo.New()
	mux.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	mux.Logger.Fatal(mux.Start(":1323"))

}
