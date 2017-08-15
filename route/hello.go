package route

import (
	"github.com/cy-kosuke-ono/goserve/action"
	"github.com/labstack/echo"
)

func HelloRouter(e *echo.Echo, a action.Action) {
	hello := e.Group("/hello")
	methods := []string{"GET", "HEAD"}

	hello.Match(methods, "/", a.Hello())
	hello.Match(methods, "/:name", a.HelloWithName())
	hello.Match(methods, "/:name/json", a.HelloWithNameJSON())
}
