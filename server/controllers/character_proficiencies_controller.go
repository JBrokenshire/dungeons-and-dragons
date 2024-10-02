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

func (c *CharacterProficienciesController) GetCharacterProficientTools(ctx echo.Context) error {
	characterProficientTools, err := c.CharacterProficienciesStore.GetProficientToolsByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	var tools []string
	for _, proficientTool := range characterProficientTools {
		tools = append(tools, proficientTool.Tool)
	}

	return ctx.JSON(http.StatusOK, tools)
}

func (c *CharacterProficienciesController) GetCharacterLanguages(ctx echo.Context) error {
	characterLanguages, err := c.CharacterProficienciesStore.GetLanguagesByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	var languages = []string{"Common"}
	for _, language := range characterLanguages {
		languages = append(languages, language.Language)
	}

	return ctx.JSON(http.StatusOK, languages)
}
