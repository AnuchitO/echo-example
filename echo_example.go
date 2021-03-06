package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/middleware"
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

	port := os.Getenv("PORT")
	e.Start(port) // listen and serve on 127.0.0.0:8080  // HL
}
