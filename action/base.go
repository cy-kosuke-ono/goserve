package action

import (
	"net/http"

	"github.com/labstack/echo"
)

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
