package action

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cy-kosuke-ono/goserve/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type HttpTest struct {
	e       *echo.Echo
	content interface{}
	h       echo.HandlerFunc
	req     *http.Request
	rec     *httptest.ResponseRecorder
	ctx     echo.Context
}

func NewHttpGetTest() HttpTest {
	h := HttpTest{
		e:       echo.New(),
		content: "",
		h:       New().Hello(),
		req:     httptest.NewRequest(echo.GET, "/", nil),
		rec:     httptest.NewRecorder(),
	}
	h.ctx = h.e.NewContext(h.req, h.rec)
	return h
}

func (h *HttpTest) UpdateContext() {
	h.ctx = h.e.NewContext(h.req, h.rec)
}

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

