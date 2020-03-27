package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/code_storage/bolt"
	"github.com/code_storage/core"
	"github.com/labstack/echo"
)

var (
	testBucketName = []byte("tmp")
	testPath       = "tmp/TemporaryStorage.db"

	badJson1 = `{
		"language": "test lang"
	}`

	badJson2 = `{
		"language": "test lang",
		"source": "test code",
		"unavailable": "test"
	}`

	badJson3 = `{
		"language": "test lang",
		"source": "    	"
	}`
)

func TestGetAddWithBadRequest(t *testing.T) {
	cs, _ := bolt.NewCodeStorage(testBucketName, testPath)
	defer cs.Close()
	h := Handler{Storage: cs}

	e := echo.New()

	req1 := httptest.NewRequest(http.MethodPost, "/codes", strings.NewReader(badJson1))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c := e.NewContext(req1, rec1)

	expectedError := echo.NewHTTPError(http.StatusBadRequest, core.ErrEmptySource.Error())
	if err := h.Add(c); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("handler didn't handle json with empty field properly and have got: %v, when expected: %v", err, expectedError)
	}

	req2 := httptest.NewRequest(http.MethodPost, "/codes", strings.NewReader(badJson2))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c = e.NewContext(req2, rec2)

	if err := h.Add(c); err != nil {
		t.Errorf("handler didn't handle json with additional field properly and have error: %v", err)
	}

	req3 := httptest.NewRequest(http.MethodPost, "/codes", strings.NewReader(badJson3))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c = e.NewContext(req3, rec3)

	if err := h.Add(c); err == nil || err.Error() == expectedError.Error() {
		t.Errorf("handler didn't handle json with only blank line source field properly and have error: %v, but expected: %v", err, expectedError)
	}
}
