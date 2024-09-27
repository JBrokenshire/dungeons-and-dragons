package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterSkillsController struct {
	CharacterSkillsStore stores.CharacterSkillsStore
}

func (c *CharacterSkillsController) GetProficientByCharacterID(ctx echo.Context) error {
	characterProficientSkills, err := c.CharacterSkillsStore.GetProficientByCharacterID(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, characterProficientSkills)
}
