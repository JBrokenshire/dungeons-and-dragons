package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterInventoryController struct {
	Store stores.CharacterInventoryStore
}

func (c *CharacterInventoryController) GetCharacterInventory(ctx echo.Context) error {
	characterInventory, err := c.Store.GetInventoryByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterInventory)
}

func (c *CharacterInventoryController) GetCharacterEquippedWeapons(ctx echo.Context) error {
	characterEquippedWeapons, err := c.Store.GetEquippedWeaponsByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterEquippedWeapons)
}
