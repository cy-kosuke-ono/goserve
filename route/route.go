package route

import (
	"github.com/cy-kosuke-ono/goserve/action"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	render := action.New()
	defer RouteRouter(e, render)

	RootRouter(e, render)
	HelloRouter(e, render)
	TrivialRouter(e, render)
}

func RouteRouter(e *echo.Echo, a action.Action) {
	methods := []string{"GET", "HEAD"}
	e.Match(methods, "/route", render.ToJSON(e.Routes()))
}
