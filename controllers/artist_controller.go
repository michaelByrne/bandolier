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

type CreateArtistRequest struct {
	Name string `json:"name"`
}
