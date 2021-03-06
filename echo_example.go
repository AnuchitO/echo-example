package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error { // HL
	return c.JSON(http.StatusOK, map[string]string{ // HL
		"message": "hello",
	}) // HL
} // HL

func main() {
	e := echo.New() // HL
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/hello", helloHandler) // HL
	e.Start(":1323")              // listen and serve on 127.0.0.0:8080  // HL
}
