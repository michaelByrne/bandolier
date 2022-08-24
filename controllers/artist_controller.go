package controllers

import (
	"bandolier/domain/showartist/commands"
	"bandolier/infrastructure"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ArtistController struct {
	dispatcher *infrastructure.Dispatcher
	eventStore infrastructure.EventStore
}

func NewArtistController(dispatcher *infrastructure.Dispatcher, es infrastructure.EventStore) *ArtistController {
	return &ArtistController{
		dispatcher: dispatcher,
		eventStore: es,
	}
}

func (c *ArtistController) Register(e *echo.Echo) {
	e.POST("/artist/create", c.CreateHandler)
	e.POST("artist/pay", c.PayHandler)
}

func (c *ArtistController) CreateHandler(ctx echo.Context) error {
	createArtistRequest := &CreateArtistRequest{}
	err := ctx.Bind(createArtistRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	command := commands.CreateArtist{
		Name: createArtistRequest.Name,
	}
	err = c.dispatcher.Dispatch(command, infrastructure.CommandMetadata{})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (c *ArtistController) PayHandler(ctx echo.Context) error {
	payArtistRequest := &PayArtistRequest{}
	err := ctx.Bind(payArtistRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	command := commands.PayArtist{
		ShowID:        payArtistRequest.ShowID,
		ArtistID:      payArtistRequest.ArtistID,
		AmountInCents: payArtistRequest.AmountInCents,
	}
	err = c.dispatcher.Dispatch(command, infrastructure.CommandMetadata{})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

type CreateArtistRequest struct {
	Name string `json:"name"`
}

type PayArtistRequest struct {
	ShowID        string `json:"show_id"`
	ArtistID      string `json:"artist_id"`
	AmountInCents int    `json:"amount_in_cents"`
}
