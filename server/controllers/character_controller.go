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
	CharacterStore stores.CharacterStore
	ClassStore     stores.ClassStore
	RaceStore      stores.RaceStore
}

func (c *CharacterController) Create(ctx echo.Context) error {
	requestCharacter := new(requests.CharacterRequest)

	// Bind new character from request body
	if err := ctx.Bind(&requestCharacter); err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}
	newCharacter, err := c.validateCharacterRequest(requestCharacter)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	// Create new character in the character stores
	err = c.CharacterStore.Create(newCharacter)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	// Return 201 Status Created and the character
	return ctx.JSON(http.StatusCreated, newCharacter)
}

func (c *CharacterController) GetAll(ctx echo.Context) error {
	characters, err := c.CharacterStore.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, characters)
}

func (c *CharacterController) Get(ctx echo.Context) error {
	// Get character using that ID
	character, err := c.CharacterStore.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, character)
}

func (c *CharacterController) Update(ctx echo.Context) error {
	existingCharacter, err := c.CharacterStore.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	// Create new models to hold the updated character
	updatedCharacterRequest := new(requests.CharacterRequest)
	// Bind the new models to the request body
	if err := ctx.Bind(&updatedCharacterRequest); err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	// Request body is empty
	if updatedCharacterRequest == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid character request body"))
	}

	// Request body contains a Character with no fields assigned
	if updatedCharacterRequest.IsEmpty() {
		return ctx.JSON(http.StatusOK, existingCharacter)
	}

	updatedCharacter, err := c.validateCharacterRequest(updatedCharacterRequest)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	updatedCharacter.ID = existingCharacter.ID

	// Update the existing character in the stores with the updated information
	err = c.CharacterStore.Update(updatedCharacter)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, updatedCharacter)
}

func (c *CharacterController) LevelUp(ctx echo.Context) error {
	character, err := c.CharacterStore.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	character.Level++

	err = c.CharacterStore.Update(character)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, character)
}

func (c *CharacterController) Delete(ctx echo.Context) error {
	err := c.CharacterStore.Delete(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, "character successfully deleted")
}

func (c *CharacterController) validateCharacterRequest(request *requests.CharacterRequest) (*models.Character, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	character := new(models.Character)
	if request.Level < 1 || request.Level > 20 {
		return nil, errors.New("invalid character level")
	}
	if !c.ClassStore.IsValidID(request.ClassID) {
		return nil, errors.New("invalid character classID")
	}
	if !c.RaceStore.IsValidID(request.RaceID) {
		return nil, errors.New("invalid character raceID")
	}

	character.Name = request.Name
	character.Level = request.Level
	character.ClassID = request.ClassID
	character.RaceID = request.RaceID

	return character, nil
}
