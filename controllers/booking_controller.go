package controllers

import (
	"bandolier/domain/readmodel"
	"bandolier/domain/venueshow"
	"bandolier/domain/venueshow/commands"
	"bandolier/infrastructure"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type BookingController struct {
	availableSlotsRepository readmodel.AvailableSlotsRepository
	showDetailRepository     readmodel.ShowDetailRepository
	dispatcher               *infrastructure.Dispatcher
	eventStore               infrastructure.EventStore
}

func NewBookingController(repo readmodel.AvailableSlotsRepository, detailRepo readmodel.ShowDetailRepository, dispatcher *infrastructure.Dispatcher, es infrastructure.EventStore) *BookingController {
	return &BookingController{
		availableSlotsRepository: repo,
		dispatcher:               dispatcher,
		eventStore:               es,
		showDetailRepository:     detailRepo,
	}
}

func (c *BookingController) Register(e *echo.Echo) {
	e.GET("/slots/available/:date", c.AvailableHandler)
	e.GET("/shows/:date", c.ShowsHandler)
	e.GET("/show/:date/:venueID", c.ShowDetailHandler)
	e.POST("/slots/schedule", c.ScheduleSlotHandler)
	e.POST("/shows/schedule", c.ScheduleShowHandler)
	e.POST("/slots/book", c.BookSlotHandler)
}

func (c *BookingController) AvailableHandler(ctx echo.Context) error {
	date := ctx.Param("date")

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	slots, err := c.availableSlotsRepository.GetSlotsAvailableOn(parsedDate)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, slots)
}

func (c *BookingController) ShowsHandler(ctx echo.Context) error {
	return nil
}

func (c *BookingController) ScheduleSlotHandler(ctx echo.Context) error {
	scheduleSlotRequest := &ScheduleSlotRequest{}
	err := ctx.Bind(scheduleSlotRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	date, err := time.Parse("Jan 2, 2006 3:04pm (MST)", scheduleSlotRequest.Date)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	duration, err := time.ParseDuration(scheduleSlotRequest.Duration)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	command := commands.NewScheduleSlot(scheduleSlotRequest.ID, date, duration, scheduleSlotRequest.VenueID)
	err = c.dispatcher.Dispatch(command, infrastructure.CommandMetadata{})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (c *BookingController) ScheduleShowHandler(ctx echo.Context) error {
	scheduleShowRequest := &ScheduleShowRequest{}
	err := ctx.Bind(scheduleShowRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	date, err := time.Parse("2006-01-02", scheduleShowRequest.Date)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	command := commands.NewScheduleShow(scheduleShowRequest.VenueID, scheduleShowRequest.VenueName, date, []commands.ScheduledSlot{})
	err = c.dispatcher.Dispatch(command, infrastructure.CommandMetadata{})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (c *BookingController) BookSlotHandler(ctx echo.Context) error {
	bookSlotRequest := &BookSlotRequest{}
	err := ctx.Bind(bookSlotRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	showID, err := c.showDetailRepository.GetShowIDForSlot(bookSlotRequest.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	command := commands.NewBookSlot(
		bookSlotRequest.ID,
		bookSlotRequest.ArtistID,
		showID,
		bookSlotRequest.ArtistName,
		bookSlotRequest.Headliner,
	)
	err = c.dispatcher.Dispatch(command, infrastructure.CommandMetadata{})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (c *BookingController) ShowDetailHandler(ctx echo.Context) error {
	date := ctx.Param("date")
	venueID := ctx.Param("venueID")

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	showDetail, err := c.showDetailRepository.GetShowDetail(venueshow.NewShowID(venueID, parsedDate).Value)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, showDetail)
}

type ScheduleShowRequest struct {
	Date      string `json:"date"`
	VenueID   string `json:"venueId"`
	VenueName string `json:"venueName"`
}

type ScheduleSlotRequest struct {
	ID       string `json:"id"`
	Date     string `json:"date"`
	Duration string `json:"duration"`
	VenueID  string `json:"venueId"`
}

type BookSlotRequest struct {
	ID         string `json:"id"`
	ArtistID   string `json:"artistId"`
	ArtistName string `json:"artistName"`
	VenueID    string `json:"venueId"`
	Headliner  bool   `json:"headliner"`
	Start      string `json:"start"`
}
