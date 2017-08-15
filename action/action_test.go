package action

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
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
