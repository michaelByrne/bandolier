package controllers

import (
	"bandolier/domain/showbank/commands"
	"bandolier/infrastructure"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BankController struct {
	dispatcher *infrastructure.Dispatcher
	eventStore infrastructure.EventStore
}

func NewBankController(dispatcher *infrastructure.Dispatcher, es infrastructure.EventStore) *BankController {
	return &BankController{
		dispatcher: dispatcher,
		eventStore: es,
	}
}

func (c *BankController) Register(e *echo.Echo) {
	e.POST("/bank/open", c.OpenHandler)
	e.POST("/bank/pay-door", c.PayDoorHandler)
	e.POST("/bank/receive-covers", c.ReceiveCoversHandler)
}

func (c *BankController) OpenHandler(ctx echo.Context) error {
	openBankRequest := &OpenBankRequest{}
	err := ctx.Bind(openBankRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	command := commands.OpenBank{
		ShowID:         openBankRequest.ShowID,
		PresaleInCents: openBankRequest.PresaleInCents,
	}
	err = c.dispatcher.Dispatch(command, infrastructure.CommandMetadata{})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (c *BankController) PayDoorHandler(ctx echo.Context) error {
	payDoorRequest := &PayDoorRequest{}
	err := ctx.Bind(payDoorRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	command := commands.PayDoor{
		ShowID:        payDoorRequest.ShowID,
		AmountInCents: payDoorRequest.AmountInCents,
	}
	err = c.dispatcher.Dispatch(command, infrastructure.CommandMetadata{})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (c *BankController) ReceiveCoversHandler(ctx echo.Context) error {
	receiveCoversRequest := &ReceiveCoversRequest{}
	err := ctx.Bind(receiveCoversRequest)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	command := commands.ReceiveCovers{
		ShowID:        receiveCoversRequest.ShowID,
		AmountInCents: receiveCoversRequest.AmountInCents,
	}
	err = c.dispatcher.Dispatch(command, infrastructure.CommandMetadata{})
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return nil
}

type OpenBankRequest struct {
	ShowID         string `json:"showID"`
	PresaleInCents int    `json:"presaleInCents"`
}

type PayDoorRequest struct {
	ShowID        string `json:"showID"`
	AmountInCents int    `json:"amountInCents"`
}

type ReceiveCoversRequest struct {
	ShowID        string `json:"showID"`
	AmountInCents int    `json:"amountInCents"`
}
