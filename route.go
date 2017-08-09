package main

import (
	"github.com/cy-kosuke-ono/goserve/action"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	render := action.New()

	e.GET("/", render.Hello())
	e.HEAD("/", render.Hello())

	e.GET("/hello", render.Hello())
	e.HEAD("/hello", render.Hello())

	e.GET("/hello/:name", render.HelloWithName())
	e.HEAD("/hello/:name", render.HelloWithName())

	e.GET("/hello/:name/json", render.HelloWithNameJSON())
	e.HEAD("/hello/:name/json", render.HelloWithNameJSON())

	e.GET("/teapod", render.Teapod())
	e.HEAD("/teapod", render.Teapod())

	e.GET("/route", render.ToJSON(e.Routes()))
	e.HEAD("/route", render.ToJSON(e.Routes()))
}
