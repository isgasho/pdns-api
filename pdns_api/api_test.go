package pdns_api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func dummyContext(t *testing.T, reqType, reqPath string, args interface{}) (echo.Context, *httptest.ResponseRecorder) {
	var rp []byte
	var jsonerr error
	if args != nil {
		switch args.(type) {
		case string:
			rp = []byte(args.(string))
		default:
			rp, jsonerr = json.Marshal(args)
			assert.NoError(t, jsonerr)
		}
	}
	req := httptest.NewRequest(reqType, reqPath, bytes.NewReader(rp))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(req, rec)

	return ctx, rec
}
