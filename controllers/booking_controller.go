package controllers

import "github.com/labstack/echo/v4"

type BookingController struct{}

func NewBookingController() *BookingController {
	return &BookingController{}
}

func (c *BookingController) Register(e *echo.Echo) {
	e.GET("/slots/available/:date", c.AvailableHandler)
	e.GET("/shows/:date", c.ShowsHandler)
	e.POST("/slots/book", c.BookSlotHandler)
	e.POST("/shows/schedule", c.ScheduleShowHandler)
}

func (c *BookingController) AvailableHandler(ctx echo.Context) error {
	return nil
}

func (c *BookingController) ShowsHandler(ctx echo.Context) error {
	return nil
}

func (c *BookingController) BookSlotHandler(ctx echo.Context) error {
	return nil
}

func (c *BookingController) ScheduleShowHandler(c2 echo.Context) error {
	return nil
}
