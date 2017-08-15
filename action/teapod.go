package action

import (
	"net/http"

	"github.com/cy-kosuke-ono/goserve/model"
	"github.com/labstack/echo"
)

func (a Action) Teapod() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusTeapot, model.TeaPod)
	}
}
