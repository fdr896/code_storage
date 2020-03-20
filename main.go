package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Code struct {
	ID       string    `json:"id,omitempty"`
	Source   string    `json:"source,omitempty"`
	Language string    `json:"language,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}

var codes = map[string]Code{
	"123": Code{ID: "123", Source: "Hello world", Language: "python", Date: time.Now()},
	"231": Code{ID: "231", Source: "print('Anton')", Language: "python", Date: time.Now()},
}

func main() {
	e := echo.New()

	e.GET("/codes/:id", Get)
	e.GET("/codes", GetAll)
	e.POST("/codes/:id", Save)

	e.Logger.Fatal(e.Start(":7070"))
}

func Get(c echo.Context) error {
	if code, ok := codes[c.Param("id")]; ok {
		return c.JSON(http.StatusOK, code)
	}
	return c.String(http.StatusNotFound, "Code with corresponding id not found")
}

func GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, codes)
}

func Save(c echo.Context) error {
	code := new(Code)
	if err := c.Bind(code); err != nil {
		return err
	}

	codes[c.Param("id")] = *code

	return c.JSON(http.StatusOK, code)
}
