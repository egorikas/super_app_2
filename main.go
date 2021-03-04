package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Привет Боб!")
	})
	i := 0
	e.GET("/second", func(c echo.Context) error {
		i++
		if i % 2 == 0{
			log.Error("этот запрос упал")
			return c.String(http.StatusBadRequest, "Боб, уйди!")
		}
		return c.String(http.StatusOK, "Боб, опять ты!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}