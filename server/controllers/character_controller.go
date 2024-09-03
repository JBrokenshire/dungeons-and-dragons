package controllers

import (
	"dungeons-and-dragons/db/models"
	"dungeons-and-dragons/db/stores"
	res "dungeons-and-dragons/server/responses"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterController struct {
	Store stores.CharacterStore
}

func (c *CharacterController) Create(ctx echo.Context) error {
	newCharacter := new(models.Character)

	// Bind new character from request body
	if err := ctx.Bind(&newCharacter); err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	if newCharacter == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid request body"))
	}

	// Create new character in the character stores
	err := c.Store.Create(newCharacter)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	// Return 201 Status Created and the character
	return ctx.JSON(http.StatusCreated, newCharacter)
}

func (c *CharacterController) GetAll(ctx echo.Context) error {
	characters, err := c.Store.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, characters)
}

func (c *CharacterController) Get(ctx echo.Context) error {
	// Get character using that ID
	character, err := c.Store.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, character)
}

func (c *CharacterController) Update(ctx echo.Context) error {
	// Create new models to hold the updated character
	updatedCharacter := new(models.Character)
	// Bind the new models to the request body
	if err := ctx.Bind(&updatedCharacter); err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}
	if updatedCharacter == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid request body"))
	}

	existingCharacter, err := c.Store.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	updatedCharacter.ID = existingCharacter.ID
	if updatedCharacter.Name == "" {
		updatedCharacter.Name = existingCharacter.Name
	}
	if updatedCharacter.Level == 0 {
		updatedCharacter.Level = existingCharacter.Level
	}
	if updatedCharacter.ClassID == 0 {
		updatedCharacter.ClassID = existingCharacter.ClassID
	}
	if updatedCharacter.RaceID == 0 {
		updatedCharacter.RaceID = existingCharacter.RaceID
	}

	// Update the existing character in the stores with the updated information
	err = c.Store.Update(updatedCharacter)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return c.Get(ctx)
}

func (c *CharacterController) LevelUp(ctx echo.Context) error {
	character, err := c.Store.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	character.Level++

	err = c.Store.Update(character)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, character)
}

func (c *CharacterController) Delete(ctx echo.Context) error {
	err := c.Store.Delete(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, "character successfully deleted")
}
