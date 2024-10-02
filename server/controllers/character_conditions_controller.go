package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterConditionsController struct {
	CharacterConditionsStore stores.CharacterConditionsStore
}

func (c *CharacterConditionsController) GetCharacterConditions(ctx echo.Context) error {
	characterConditions, err := c.CharacterConditionsStore.GetConditionsByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterConditions)
}
