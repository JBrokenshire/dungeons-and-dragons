package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterDefensesController struct {
	CharacterDefensesStore stores.CharacterDefensesStore
}

func (c *CharacterDefensesController) GetCharacterDefenses(ctx echo.Context) error {
	characterDefenses, err := c.CharacterDefensesStore.GetDefensesByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterDefenses)
}
