package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterProficienciesController struct {
	CharacterProficienciesStore stores.CharacterProficienciesStore
}

func (c *CharacterProficienciesController) GetCharacterProficientArmourTypes(ctx echo.Context) error {
	characterProficientArmourTypes, err := c.CharacterProficienciesStore.GetProficientArmourTypesByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	var armourTypes []string
	for _, proficientArmourType := range characterProficientArmourTypes {
		armourTypes = append(armourTypes, proficientArmourType.ArmourType)
	}

	return ctx.JSON(http.StatusOK, armourTypes)
}

func (c *CharacterProficienciesController) GetCharacterProficientWeapons(ctx echo.Context) error {
	characterProficientWeapons, err := c.CharacterProficienciesStore.GetProficientWeaponsByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	var weapons []string
	for _, proficientWeapon := range characterProficientWeapons {
		weapons = append(weapons, proficientWeapon.Weapon)
	}

	return ctx.JSON(http.StatusOK, weapons)
}
