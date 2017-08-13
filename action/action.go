package action

import (
	"fmt"
	"net/http"

	"github.com/cy-kosuke-ono/goserve/model"
	"github.com/labstack/echo"
)

type Action struct{}

func New() Action {
	return Action{}
}

func (a Action) ToPlain(i interface{}) echo.HandlerFunc {
	switch str := i.(type) {
	case string:
		return func(c echo.Context) error {
			return c.String(http.StatusOK, str)
		}
	default:
		return func(c echo.Context) error {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				"Cannot convert to string.",
			)
		}
	}
}

func (a Action) ToJSON(i interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, i, "  ")
	}
}

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

func (a Action) Teapod() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusTeapot, model.TeaPod)
	}
}
