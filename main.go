package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rest-api/core"
	"github.com/rest-api/memory"
)

var storage memory.CodeStorage

func main() {
	e := echo.New()

	e.GET("/codes/:id", Get)
	e.GET("/codes", GetAll)
	e.POST("/codes", Add)

	e.Logger.Fatal(e.Start(":8080"))
}

// Get function handles GET request (return code by its id)
func Get(c echo.Context) error {
	code, err := storage.Get(c.Param("id"))

	if err != nil {
		return err
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

	if err := storage.AddCode(code); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, code)
}
