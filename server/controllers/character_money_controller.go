package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterMoneyController struct {
	Store stores.CharacterMoneyStore
}

func (c *CharacterMoneyController) GetCharacterMoney(ctx echo.Context) error {
	characterMoney, err := c.Store.GetMoneyByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterMoney)
}
