package main

import (
	"log"

	"github.com/fdr896/code_storage/rest-api/bolt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	bucketsNames = [][]byte{
		[]byte("code"),
		[]byte("user"),
	}
	path = "codestorage.db"
)

func main() {
	h := &Handler{bolt.New(path, bucketsNames)}

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/codes", h.GetAll)
	e.GET("/codes/:id", h.Get)
	e.GET("/user/:login", h.IsNewUser)

	e.POST("/codes", h.Add)
	e.POST("/user", h.AddUser)
	e.POST("/user/:login", h.VerifyPassword)
	e.POST("/user/:login/set", h.SetPassword)

	e.DELETE("/codes/:id", h.Delete)

	log.Fatal(e.Start(":8080"))
}
