package main

import (
	"github.com/cy-kosuke-ono/goserve/action"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	render := action.New()

	methods := []string{"GET", "HEAD"}

	e.Match(methods, "/", render.Hello())
	e.Match(methods, "/hello", render.Hello())
	e.Match(methods, "/hello/:name", render.HelloWithName())
	e.Match(methods, "/hello/:name/json", render.HelloWithNameJSON())
	e.Match(methods, "/teapod", render.Teapod())
	e.Match(methods, "/route", render.ToJSON(e.Routes()))
}
