package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())

	e.GET("/", handler)

	e.Logger.Fatal(e.Start(":8080"))
}

func handler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}
