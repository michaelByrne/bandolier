package main

import (
	"bandolier/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	println("Hello, world!")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bookingController := controllers.NewBookingController()
	bookingController.Register(e)

	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":5001"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Bandolier!")
}
