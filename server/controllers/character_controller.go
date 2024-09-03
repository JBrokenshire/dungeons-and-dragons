package controllers

import (
	"dungeons-and-dragons/db/models"
	"dungeons-and-dragons/db/stores"
	"dungeons-and-dragons/server/requests"
	res "dungeons-and-dragons/server/responses"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CharacterController struct {
	Store stores.CharacterStore
}

func (c *CharacterController) Create(ctx echo.Context) error {
	requestCharacter := new(requests.CharacterRequest)

	// Bind new character from request body
	if err := ctx.Bind(&requestCharacter); err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}
	if requestCharacter == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid request body"))
	}

	if !c.Store.IsValidClassID(requestCharacter.ClassID) {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid classID"))
	}
	if !c.Store.IsValidRaceID(requestCharacter.RaceID) {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid raceID"))
	}
	if requestCharacter.Level < 1 || requestCharacter.Level > 20 {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid level"))
	}

	newCharacter := &models.Character{
		Name:    requestCharacter.Name,
		Level:   requestCharacter.Level,
		ClassID: requestCharacter.ClassID,
		RaceID:  requestCharacter.RaceID,
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
	updatedCharacterRequest := new(requests.CharacterRequest)
	// Bind the new models to the request body
	if err := ctx.Bind(&updatedCharacterRequest); err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}
	if updatedCharacterRequest == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid request body"))
	}

	existingCharacter, err := c.Store.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	updatedCharacter := &models.Character{
		ID:      existingCharacter.ID,
		Name:    existingCharacter.Name,
		Level:   existingCharacter.Level,
		ClassID: existingCharacter.ClassID,
		RaceID:  existingCharacter.RaceID,
	}
	if updatedCharacterRequest.Name != "" {
		updatedCharacter.Name = updatedCharacterRequest.Name
	}
	if updatedCharacterRequest.Level != 0 {
		updatedCharacter.Level = updatedCharacterRequest.Level
	}
	if updatedCharacterRequest.ClassID != 0 {
		if !c.Store.IsValidClassID(updatedCharacterRequest.ClassID) {
			return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid classID"))
		}
		updatedCharacter.ClassID = updatedCharacterRequest.ClassID
	}
	if updatedCharacterRequest.RaceID != 0 {

		if !c.Store.IsValidRaceID(updatedCharacterRequest.RaceID) {
			return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid raceID"))
		}
		updatedCharacter.RaceID = updatedCharacterRequest.RaceID
	}

	// Update the existing character in the stores with the updated information
	err = c.Store.Update(updatedCharacter)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, updatedCharacter)
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
