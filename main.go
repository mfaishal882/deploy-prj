package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	log.Print("\n\n\n")
	e.Use(middleware.Logger())
	log.Print("\n\n\n")
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	e.GET("/hello_gaes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello gaes from wonderland")
	})
	e.GET("/hello_bang", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello bang makan bang")
	})
	e.GET("/hello_coy", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello coy makan kuy")
	})
	e.Start(":8000")
}
