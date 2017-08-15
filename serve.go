package main

import (
	"github.com/cy-kosuke-ono/goserve/base"
	"github.com/cy-kosuke-ono/goserve/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	base := base.New()
	defer base.CloseLogFile()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: base.LogFileFd,
	}))

	route.Router(e)

	e.Logger.Fatal(e.Start(":" + *base.Port))
}
