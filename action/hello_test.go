package action

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cy-kosuke-ono/goserve/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	c := NewHttpGetTest()
	c.h = New().Hello()
	c.req.Header.Set(
		echo.HeaderContentType,
		echo.MIMETextPlain,
	)
	c.UpdateContext()

	if assert.NoError(t, c.h(c.ctx)) {
		assert.Equal(t, http.StatusOK, c.rec.Code)
		assert.Equal(t, "Hello, world.", c.rec.Body.String())
	}
}

func TestHelloWithNameHandler(t *testing.T) {
	name := "bogus"
	c := NewHttpGetTest()
	c.h = New().HelloWithName()
	c.req.Header.Set(
		echo.HeaderContentType,
		echo.MIMETextPlain,
	)

	// Build context. Depend on order.
	c.UpdateContext()
	c.ctx.SetParamNames("name")
	c.ctx.SetParamValues(name)

	if assert.NoError(t, c.h(c.ctx)) {
		assert.Equal(t, http.StatusOK, c.rec.Code)
		assert.Equal(t, "Hello, "+name, c.rec.Body.String())
	}
}

func TestHelloWithNameJSONHandler(t *testing.T) {
	name := "phony"
	expected := model.Person{name, http.StatusOK}
	actual := model.Person{}
	assert := assert.New(t)

	c := NewHttpGetTest()
	c.h = New().HelloWithNameJSON()
	c.req.Header.Set(
		echo.HeaderContentType,
		echo.MIMEApplicationJSON,
	)

	// Build context. Depend on order.
	c.UpdateContext()
	c.ctx.SetParamNames("name")
	c.ctx.SetParamValues(name)

	if assert.NoError(c.h(c.ctx)) {
		assert.Equal(http.StatusOK, c.rec.Code)
		if j, err := json.Marshal(expected); err == nil {
			assert.NotEmpty(c.rec.Body.Bytes())
			assert.JSONEq(string(j), c.rec.Body.String())
			assert.Contains(c.rec.Body.String(), name)
		}
		assert.NoError(json.Unmarshal(c.rec.Body.Bytes(), &actual))
		assert.Equal(expected, actual)
	}
}
