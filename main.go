package main

import (
	"log"
	"net/http"

	"github.com/code_storage/bolt"
	"github.com/code_storage/core"
	"github.com/labstack/echo"
)

var (
	bucketName = []byte("CodeStorage")
	path       = "codestorage.db"
)

func main() {
	var h Handler
	if storage, err := bolt.NewCodeStorage(bucketName, path); err != nil {
		log.Fatal(err)
	} else {
		h.Storage = storage
	}
	defer h.Storage.Close()

	e := echo.New()
	e.GET("/codes/:id", h.Get)
	e.GET("/codes", h.GetAll)
	e.POST("/codes", h.Add)
	e.DELETE("/codes/:id", h.Delete)

	log.Fatal(e.Start(":8080"))
}

// Handler handles http request.
type Handler struct {
	Storage core.CodeStorage
}

// Get handles GET request (return code by its id).
func (h *Handler) Get(c echo.Context) error {
	code, err := h.Storage.Get(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, core.ErrNotFound)
	}

	return c.JSON(http.StatusOK, code)
}

// GetAll handles GET request (return all codes).
func (h *Handler) GetAll(c echo.Context) error {
	codes, err := h.Storage.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, codes)
}

// Add adds received code snippet to database.
func (h *Handler) Add(c echo.Context) error {
	code := &core.Code{}

	if err := c.Bind(code); err != nil {
		return err
	}

	if err := code.NewCode(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.Storage.Add(code); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, code)
}

// Delete deletes code snippet from database by its id.
func (h *Handler) Delete(c echo.Context) error {
	if err := h.Storage.Delete(c.Param("id")); err != nil {
		return err
	}
	return nil
}
