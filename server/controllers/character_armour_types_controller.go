package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterArmourTypesController struct {
	CharacterArmourTypesStore stores.CharacterArmourTypesStore
}

func (c *CharacterArmourTypesController) GetProficientArmourTypes(ctx echo.Context) error {
	characterProficientArmourTypes, err := c.CharacterArmourTypesStore.GetProficientArmourTypesByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterProficientArmourTypes)
}
