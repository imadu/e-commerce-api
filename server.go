package main

import (
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/imadu/e-commerce-api/db"
	"github.com/labstack/echo"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("Could not load env file: %+v", err)
	}

	db.InitDB()

	mux := echo.New()
	mux.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	mux.Logger.Fatal(mux.Start(":1323"))

}
