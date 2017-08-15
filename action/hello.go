package action

import (
	"fmt"
	"net/http"

	"github.com/cy-kosuke-ono/goserve/model"
	"github.com/labstack/echo"
)

func (a Action) Hello() echo.HandlerFunc {
	return a.ToPlain("Hello, world.")
}

func (a Action) HelloWithName() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(
			http.StatusOK,
			fmt.Sprintf("Hello, %v", c.Param("name")),
		)
	}
}

func (a Action) HelloWithNameJSON() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSONPretty(
			http.StatusOK,
			model.Person{c.Param("name"), http.StatusOK},
			"  ",
		)
	}
}
