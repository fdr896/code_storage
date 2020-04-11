package main

import (
	"net/http"

	"github.com/fdr896/code_storage/rest-api/bolt"
	"github.com/fdr896/code_storage/rest-api/core"
	"github.com/labstack/echo"
)

// Handler handles http requests.
type Handler struct {
	Storage *bolt.Bolt
}

// Get handles GET request (return code by its id).
func (h *Handler) Get(c echo.Context) error {
	code, err := h.Storage.CodeStorage.Get(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, core.ErrNotFound)
	}

	return c.JSONPretty(http.StatusOK, code, "\t")
}

// GetAll handles GET request (return all codes).
func (h *Handler) GetAll(c echo.Context) error {
	codes, err := h.Storage.CodeStorage.GetAll()
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusOK, codes, "\t")
}

// Add adds received code snippet to database.
func (h *Handler) Add(c echo.Context) error {
	code := &core.Code{}

	if err := c.Bind(code); err != nil {
		return err
	}

	if err := code.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.Storage.CodeStorage.Add(code); err != nil {
		return err
	}

	return c.JSONPretty(http.StatusCreated, code, "\t")
}

// Delete deletes code snippet from database by its id.
func (h *Handler) Delete(c echo.Context) error {
	if err := h.Storage.CodeStorage.Delete(c.Param("id")); err != nil {
		return err
	}
	return nil
}

// AddUser adds user to database if there's no user with same login.
func (h *Handler) AddUser(c echo.Context) error {
	user := &core.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	if !user.Valid() {
		return echo.NewHTTPError(http.StatusBadRequest, core.ErrEmptyLogin.Error())
	}

	if err := h.Storage.UserStorage.AddUser(user); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	return c.JSONPretty(http.StatusCreated, user, "\t")
}

// VerifyPassword verifies if password with which client tries to sign in is correct password for user account with required login.
func (h *Handler) VerifyPassword(c echo.Context) error {
	login := c.Param("login")
	body := &struct {
		Password string `json:"password"`
	}{}

	if err := c.Bind(body); err != nil {
		return err
	}

	verified, err := h.Storage.UserStorage.ComparePassword(login, body.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSONPretty(http.StatusOK, verified, "\t")
}

// IsNewUser checks if user has already set password for his account.
func (h *Handler) IsNewUser(c echo.Context) error {
	login := c.Param("login")

	havePassword, err := h.Storage.UserStorage.HasPassword(login)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSONPretty(http.StatusOK, havePassword, "\t")
}

// SetPassword sets new password to account with required login if such account exists.
func (h *Handler) SetPassword(c echo.Context) error {
	login := c.Param("login")
	body := &struct {
		Password string `json:"password"`
	}{}

	if err := c.Bind(body); err != nil {
		return err
	}

	if err := h.Storage.UserStorage.SetPassword(login, body.Password); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSONPretty(http.StatusCreated, body, "\t")
}
