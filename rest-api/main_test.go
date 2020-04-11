package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/fdr896/code_storage/rest-api/bolt"
	"github.com/fdr896/code_storage/rest-api/core"
	"github.com/labstack/echo"
)

var (
	testBucketsNames = [][]byte{
		[]byte("code"),
		[]byte("user"),
	}
	testPath = "tmp/TemporaryStorageMain.db"

	badJSON1 = `{
		"language": "test lang",
		"description": "test descr"
	}`
	badJSON2 = `{
		"language": "test lang",
		"source": "test code",
		"description": "test descr",
		"unavailable": "test"
	}`
	badJSON3 = `{
		"language": "test lang",
		"source": "    ",
		"description": "test descr"
	}`

	validUser1 = `{
		"login": "admin"
	}`
	invalidUser1 = `{
		"login": "   "
	}`

	password1 = `{
		"password": "123"
	}`
	password2 = `{
		"password": "321"
	}`
	password3 = `{
		"password": ""
	}`
)

func TestGetAddWithBadRequest(t *testing.T) {
	os.RemoveAll(testPath)
	h := Handler{&bolt.Bolt{
		CodeStorage: bolt.New(testPath, testBucketsNames).CodeStorage,
	}}

	e := echo.New()

	req1 := httptest.NewRequest(http.MethodPost, "/codes", strings.NewReader(badJSON1))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res1 := httptest.NewRecorder()
	c := e.NewContext(req1, res1)

	expectedError := echo.NewHTTPError(http.StatusBadRequest, core.ErrEmptySource.Error())
	if err := h.Add(c); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("handler didn't handle json with empty field properly and have got: %v, when expected: %v", err, expectedError)
	}

	req2 := httptest.NewRequest(http.MethodPost, "/codes", strings.NewReader(badJSON2))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res2 := httptest.NewRecorder()
	c = e.NewContext(req2, res2)

	if err := h.Add(c); err != nil {
		t.Errorf("handler didn't handle json with additional field properly and have error: %v", err)
	}

	req3 := httptest.NewRequest(http.MethodPost, "/codes", strings.NewReader(badJSON3))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res3 := httptest.NewRecorder()
	c = e.NewContext(req3, res3)

	if err := h.Add(c); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("handler didn't handle json with only blank line source field properly and have error: %v, but expected: %v", err, expectedError)
	}
}

func TestAddUser(t *testing.T) {
	os.RemoveAll(testPath)
	h := &Handler{&bolt.Bolt{
		UserStorage: bolt.New(testPath, testBucketsNames).UserStorage,
	}}

	e := echo.New()

	req1 := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(validUser1))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res1 := httptest.NewRecorder()
	c := e.NewContext(req1, res1)

	if err := h.AddUser(c); err != nil {
		t.Errorf("handler didn't add user with valid login and this error occured: %v", err)
	}

	req2 := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(validUser1))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res2 := httptest.NewRecorder()
	c = e.NewContext(req2, res2)

	expectedError := echo.NewHTTPError(http.StatusForbidden, core.ErrUserWithSameLogin.Error())
	if err := h.AddUser(c); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("handler didn't handle user with already registered login properly and this error occured: %v, but expected: %v", err, expectedError)
	}

	req3 := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(invalidUser1))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res3 := httptest.NewRecorder()
	c = e.NewContext(req3, res3)

	expectedError = echo.NewHTTPError(http.StatusBadRequest, core.ErrEmptyLogin)
	if err := h.AddUser(c); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("handler doesn't handle request with empty Login field and this error occured: %v, but expected: %v", err, expectedError)
	}
}

func TestSetPassword(t *testing.T) {
	os.RemoveAll(testPath)
	h := &Handler{&bolt.Bolt{
		UserStorage: bolt.New(testPath, testBucketsNames).UserStorage,
	}}

	e := echo.New()

	req1 := httptest.NewRequest(http.MethodPost, "/user/admin/set", strings.NewReader(password1))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res1 := httptest.NewRecorder()
	c := e.NewContext(req1, res1)
	c.SetParamNames("login")
	c.SetParamValues("admin")

	expectedError := echo.NewHTTPError(http.StatusBadRequest, core.ErrNoSuchUser)
	if err := h.SetPassword(c); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("handler didn't recognise unregistered user and this error occured: %v, but expected: %v", err, expectedError)
	}

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(validUser1))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c = e.NewContext(req, res)
	h.AddUser(c)

	req2 := httptest.NewRequest(http.MethodPost, "/user/admin/set", strings.NewReader(password1))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res2 := httptest.NewRecorder()
	c = e.NewContext(req2, res2)
	c.SetParamNames("login")
	c.SetParamValues("admin")

	if err := h.SetPassword(c); err != nil {
		t.Errorf("test failed when tried to set password with this error: %v", err)
	}

	req3 := httptest.NewRequest(http.MethodPost, "/user/admin/set", strings.NewReader(password3))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res3 := httptest.NewRecorder()
	c = e.NewContext(req3, res3)
	c.SetParamNames("login")
	c.SetParamValues("admin")

	expectedError = echo.NewHTTPError(http.StatusBadRequest, core.ErrEmptyPassword.Error())
	if err := h.SetPassword(c); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("handler doesn't recognise empty password field and this error occured: %v, but expected: %v", err, expectedError)
	}
}

func TestVerifyPassword(t *testing.T) {
	os.RemoveAll(testPath)
	h := &Handler{&bolt.Bolt{
		UserStorage: bolt.New(testPath, testBucketsNames).UserStorage,
	}}

	e := echo.New()

	req1 := httptest.NewRequest(http.MethodPost, "/user/admin", strings.NewReader(password1))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res1 := httptest.NewRecorder()
	c := e.NewContext(req1, res1)
	c.SetParamNames("login")
	c.SetParamValues("admin")

	expectedError := echo.NewHTTPError(http.StatusNotFound, core.ErrNoSuchUser.Error())
	if err := h.VerifyPassword(c); err == nil || err.Error() != expectedError.Error() {
		t.Errorf("handler didn't handle not registered user and error occured: %v, but expected: %v", err, expectedError)
	}

	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(validUser1))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c = e.NewContext(req, res)
	h.AddUser(c)

	req = httptest.NewRequest(http.MethodPost, "/user/admin/set", strings.NewReader(password1))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	c.SetParamNames("login")
	c.SetParamValues("admin")
	h.SetPassword(c)

	req2 := httptest.NewRequest(http.MethodPost, "/user/admin", strings.NewReader(password1))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res2 := httptest.NewRecorder()
	c = e.NewContext(req2, res2)
	c.SetParamNames("login")
	c.SetParamValues("admin")

	if err := h.VerifyPassword(c); err != nil {
		t.Errorf("test failed when tried to verify user's password with error: %v", err)
	}

	verified, _ := strconv.ParseBool(strings.TrimSpace(res2.Body.String()))
	if !verified {
		t.Errorf("handler didn't verify correct password")
	}

	req3 := httptest.NewRequest(http.MethodPost, "/user/admin", strings.NewReader(password2))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res3 := httptest.NewRecorder()
	c = e.NewContext(req3, res3)
	c.SetParamNames("login")
	c.SetParamValues("admin")

	if err := h.VerifyPassword(c); err != nil {
		t.Errorf("test failed when tried to verify user's password with error: %v", err)
	}

	verified, _ = strconv.ParseBool(strings.TrimSpace(res3.Body.String()))
	if verified {
		t.Errorf("handler verified incorrect password")
	}
}
