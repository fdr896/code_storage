package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rest-api/core"
	"github.com/rest-api/memory"
)

var storage = memory.NewCodeStorage([]byte("CodeStorage"))

func main() {
	e := echo.New()

	e.GET("/codes/:id", Get)
	e.GET("/codes", GetAll)
	e.POST("/codes", Add)
	e.DELETE("/codes/:id", Delete)

	e.Logger.Fatal(e.Start(":8080"))
	storage.CloseDB()
}

// Get function handles GET request (return code by its id)
func Get(c echo.Context) error {
	code, err := storage.Get(c.Param("id"))

	if err != nil {
		return c.String(core.CodeDoesNotExist.StatusCode, core.CodeDoesNotExist.ErrorMessage.Error())
	}
	return c.JSON(http.StatusOK, code)
}

// GetAll function handles GET request (return all codes)
func GetAll(c echo.Context) error {
	codes, err := storage.GetAll()

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, codes)
}

// Add adds received code snippet to database
func Add(c echo.Context) error {
	code := new(core.Code)

	if err := c.Bind(code); err != nil {
		return err
	}

	if err := storage.Add(code); err != nil {
		if err.Error() == core.CodeDoesNotExist.ErrorMessage.Error() {
			return c.String(core.UnsupportedJSON.StatusCode, err.Error())
		}
		return c.String(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, code)
}

// Delete deletes code snippet from database by its id
func Delete(c echo.Context) error {
	if err := storage.Delete(c.Param("id")); err != nil {
		return c.String(core.CodeDoesNotExist.StatusCode, core.CodeDoesNotExist.ErrorMessage.Error())
	}

	return nil
}
