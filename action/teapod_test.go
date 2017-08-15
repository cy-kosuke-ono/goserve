package action

import (
	"net/http"
	"testing"

	"github.com/cy-kosuke-ono/goserve/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestTeapodHandler(t *testing.T) {
	c := NewHttpGetTest()
	c.h = New().Teapod()
	c.req.Header.Set(
		echo.HeaderContentType,
		echo.MIMETextPlain,
	)
	c.UpdateContext()

	if assert.NoError(t, c.h(c.ctx)) {
		assert.Equal(t, http.StatusTeapot, c.rec.Code)
		assert.Equal(t, model.TeaPod, c.rec.Body.String())
	}
}
