package controllers

import (
	"dnd-api/db/models"
	"dnd-api/db/stores"
	"dnd-api/server/requests"
	res "dnd-api/server/responses"
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
	if err = ctx.Bind(&updatedCharacterRequest); err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	// Request body is empty
	if updatedCharacterRequest == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid character request body"))
	}

	// None of the request body fields match the character request object
	if updatedCharacterRequest.IsEmpty() {
		return ctx.JSON(http.StatusOK, existingCharacter)
	}
	if updatedCharacterRequest.Name == "" {
		updatedCharacterRequest.Name = existingCharacter.Name
	}
	if updatedCharacterRequest.Level == 0 {
		updatedCharacterRequest.Level = existingCharacter.Level
	}
	if updatedCharacterRequest.ClassID == 0 {
		updatedCharacterRequest.ClassID = existingCharacter.ClassID
	}
	if updatedCharacterRequest.RaceID == 0 {
		updatedCharacterRequest.RaceID = existingCharacter.RaceID
	}
	if updatedCharacterRequest.Strength == 0 {
		updatedCharacterRequest.Strength = existingCharacter.Strength
	}
	if updatedCharacterRequest.Dexterity == 0 {
		updatedCharacterRequest.Dexterity = existingCharacter.Dexterity
	}
	if updatedCharacterRequest.Constitution == 0 {
		updatedCharacterRequest.Constitution = existingCharacter.Constitution
	}
	if updatedCharacterRequest.Intelligence == 0 {
		updatedCharacterRequest.Intelligence = existingCharacter.Intelligence
	}
	if updatedCharacterRequest.Wisdom == 0 {
		updatedCharacterRequest.Wisdom = existingCharacter.Wisdom
	}
	if updatedCharacterRequest.Charisma == 0 {
		updatedCharacterRequest.Charisma = existingCharacter.Charisma
	}
	if updatedCharacterRequest.WalkingSpeedModifier == 0 {
		updatedCharacterRequest.WalkingSpeedModifier = existingCharacter.WalkingSpeedModifier
	}
	if existingCharacter.Inspiration {
		updatedCharacterRequest.Inspiration = existingCharacter.Inspiration
	}
	if updatedCharacterRequest.CurrentHitPoints == 0 {
		updatedCharacterRequest.CurrentHitPoints = existingCharacter.CurrentHitPoints
	}
	if updatedCharacterRequest.MaxHitPoints == 0 {
		updatedCharacterRequest.MaxHitPoints = existingCharacter.MaxHitPoints
	}
	if updatedCharacterRequest.TempHitPoints == 0 {
		updatedCharacterRequest.TempHitPoints = existingCharacter.TempHitPoints
	}

	_, err = c.validateCharacterRequest(updatedCharacterRequest)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	existingCharacter = &models.Character{
		ID:                   existingCharacter.ID,
		Name:                 updatedCharacterRequest.Name,
		Level:                updatedCharacterRequest.Level,
		ProfilePictureURL:    updatedCharacterRequest.ProfilePictureURL,
		ClassID:              updatedCharacterRequest.ClassID,
		RaceID:               updatedCharacterRequest.RaceID,
		Strength:             updatedCharacterRequest.Strength,
		Dexterity:            updatedCharacterRequest.Dexterity,
		Constitution:         updatedCharacterRequest.Constitution,
		Intelligence:         updatedCharacterRequest.Intelligence,
		Wisdom:               updatedCharacterRequest.Wisdom,
		Charisma:             updatedCharacterRequest.Charisma,
		WalkingSpeedModifier: updatedCharacterRequest.WalkingSpeedModifier,
		Inspiration:          updatedCharacterRequest.Inspiration,
		CurrentHitPoints:     updatedCharacterRequest.CurrentHitPoints,
		MaxHitPoints:         updatedCharacterRequest.MaxHitPoints,
		TempHitPoints:        updatedCharacterRequest.TempHitPoints,
	}

	// Update the existing character in the stores with the updated information
	err = c.CharacterStore.Update(existingCharacter)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, existingCharacter)
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

func (c *CharacterController) ToggleInspiration(ctx echo.Context) error {
	character, err := c.CharacterStore.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	character.Inspiration = !character.Inspiration

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
	if request.Name == "" {
		return nil, errors.New("invalid character name")
	}
	if request.Level != 0 {
		if request.Level < 1 || request.Level > 20 {
			return nil, errors.New("invalid character level")
		}
	}
	if !c.ClassStore.IsValidID(request.ClassID) {
		return nil, errors.New("invalid character classID")
	}
	if !c.RaceStore.IsValidID(request.RaceID) {
		return nil, errors.New("invalid character raceID")
	}

	if request.MaxHitPoints <= 0 {
		return nil, errors.New("invalid character maxHitPoints")
	}
	if request.CurrentHitPoints <= 0 && request.CurrentHitPoints <= request.MaxHitPoints {
		return nil, errors.New("invalid character currentHitPoints")
	}
	if request.TempHitPoints < 0 {
		return nil, errors.New("invalid character tempHitPoints")
	}
	if request.WalkingSpeedModifier < 0 {
		return nil, errors.New("invalid character walkingSpeedModifier")
	}

	character.Name = request.Name
	character.Level = request.Level
	character.ClassID = request.ClassID
	character.RaceID = request.RaceID
	character.ProfilePictureURL = request.ProfilePictureURL
	character.Strength = request.Strength
	character.Dexterity = request.Dexterity
	character.Constitution = request.Constitution
	character.Intelligence = request.Intelligence
	character.Wisdom = request.Wisdom
	character.Charisma = request.Charisma
	character.WalkingSpeedModifier = request.WalkingSpeedModifier
	character.Inspiration = request.Inspiration
	character.CurrentHitPoints = request.CurrentHitPoints
	character.MaxHitPoints = request.MaxHitPoints
	character.TempHitPoints = request.TempHitPoints

	return character, nil
}
