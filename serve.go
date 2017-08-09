package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Person .. For Generate JSON
type Person struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func router(e *echo.Echo) {
	teapod := `
                       (
            _           ) )
         _,(_)._        ((
    ___,(_______).        )
  ,'__.   /       \    /\_
 /,' /  |""|       \  /  /
| | |   |__|       |,'  /
 \'.|                  /
   . :           :    /
     .            :.,'
       -.________,-'
`

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/hello")
	})
	e.HEAD("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/hello")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.HEAD("/hello", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/teapod", func(c echo.Context) error {
		return c.String(http.StatusTeapot, teapod)
	})
	e.HEAD("/teapod", func(c echo.Context) error {
		return c.NoContent(http.StatusTeapot)
	})

	e.GET("/hello/:name", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, "+c.Param("name"))
	})
	e.HEAD("/hello/:name", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/hello/:name/json", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, &Person{c.Param("name"), http.StatusOK}, "  ")
	})
	e.HEAD("/hello/:name/json", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/route", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, e.Routes(), "  ")
	})
	e.HEAD("/route", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}

func main() {

	e := echo.New()

	logPath := flag.String("l", "/var/log/goserve/access.log", "Access log file path")
	port := flag.String("p", "80", "Using port number")
	flag.Parse()

	fp, err := os.OpenFile(*logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fp,
	}))

	router(e)

	e.Logger.Fatal(e.Start(":" + *port))
}
