package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	log.Println("port:", port)
	log.Fatal(e.Start(":" + port))
}
