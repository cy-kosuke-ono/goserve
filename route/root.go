package route

import (
	"github.com/cy-kosuke-ono/goserve/action"
	"github.com/labstack/echo"
)

func RootRouter(e *echo.Echo, a action.Action) {
	methods := []string{"GET", "HEAD"}
	e.Match(methods, "/", a.Hello())
}

func TrivialRouter(e *echo.Echo, a action.Action) {
	methods := []string{"GET", "HEAD"}
	e.Match(methods, "/teapod", a.Teapod())
	e.Match(methods, "/error", a.ToPlain(1))
}
