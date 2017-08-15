package action

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cy-kosuke-ono/goserve/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestToPlainHandler(t *testing.T) {
	c := NewHttpGetTest()
	c.content = "hello"
	c.h = New().ToPlain(c.content)
	c.req.Header.Set(
		echo.HeaderContentType,
		echo.MIMETextPlain,
	)
	c.UpdateContext()

	if assert.NoError(t, c.h(c.ctx)) {
		assert.Equal(t, http.StatusOK, c.rec.Code)
		assert.Equal(t, c.content, c.rec.Body.String())
		assert.Contains(
			t,
			c.rec.Header().Get(echo.HeaderContentType),
			echo.MIMETextPlain,
		)
	}
}

func TestToJSONHandler(t *testing.T) {
	m := model.Person{
		Name:       "fake",
		StatusCode: http.StatusTeapot,
	}
	c := NewHttpGetTest()
	c.h = New().ToJSON(m)
	c.req.Header.Set(
		echo.HeaderContentType,
		echo.MIMEApplicationJSON,
	)
	c.UpdateContext()

	if assert.NoError(t, c.h(c.ctx)) {
		assert.Equal(t, http.StatusOK, c.rec.Code)
		if j, err := json.MarshalIndent(m, "", "  "); err == nil {
			assert.NotEmpty(t, c.rec.Body.Bytes())
			assert.JSONEq(t, string(j), c.rec.Body.String())
		}
	}
}

// TODO: echo.NewHttpError should be 500, but return 200
// func TestToPlainTypeAssertFail(t *testing.T) {
// 	c := NewHttpGetTest()
// 	c.h = New().ToPlain(1)
// 	c.req.Header.Set(
// 		echo.HeaderContentType,
// 		echo.MIMETextPlain,
// 	)
// 	c.UpdateContext()
//
// 	if assert.Error(t, c.h(c.ctx)) {
// 		assert.Equal(
// 			t,
// 			http.StatusInternalServerError,
// 			c.rec.Code,
// 		)
// 	}
// }
