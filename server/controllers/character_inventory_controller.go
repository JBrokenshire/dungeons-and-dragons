package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"errors"
	"fmt"
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

func (c *CharacterInventoryController) ToggleItemEquipped(ctx echo.Context) error {
	inventoryItem, err := c.Store.GetCharacterInventoryItemByID(ctx.Param("characterID"), ctx.Param("itemID"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	if !inventoryItem.Item.Equippable {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("item can not be equipped"))
	}

	if inventoryItem.Location != "Equipment" {
		return ctx.JSON(http.StatusOK, inventoryItem)
	}

	fmt.Println(">>> ", inventoryItem)

	inventoryItem.Equipped = !inventoryItem.Equipped
	err = c.Store.UpdateCharacterInventoryItem(inventoryItem)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, inventoryItem)
}
