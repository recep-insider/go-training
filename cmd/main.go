package main

import (
	"go-training/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/", handlers.Home)

	e.Logger.Fatal(e.Start(":1323"))
}
