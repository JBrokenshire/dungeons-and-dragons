package controller

import (
	"dungeons-and-dragons/models"
	"dungeons-and-dragons/store"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CharacterController struct {
	Store store.CharacterStore
}

func (c *CharacterController) Create(ctx echo.Context) error {
	newCharacter := new(models.Character)

	// Bind new character from request body
	if err := ctx.Bind(&newCharacter); err != nil {
		return returnErrorCode(ctx, http.StatusInternalServerError, err)
	}

	// Create new character in the character store
	err := c.Store.Create(newCharacter)
	if err != nil {
		return returnErrorCode(ctx, http.StatusInternalServerError, err)
	}

	// Return 201 Status Created and the character
	return ctx.JSON(http.StatusCreated, newCharacter)
}

func (c *CharacterController) GetAll(ctx echo.Context) error {
	characters, err := c.Store.GetAll()
	if err != nil {
		return returnErrorCode(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, characters)
}

func (c *CharacterController) Get(ctx echo.Context) error {
	// Convert ID to integer from string
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return returnErrorCode(ctx, http.StatusBadRequest, err)
	}

	// Get character using that ID
	character, err := c.Store.Get(id)
	if err != nil {
		return returnErrorCode(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, character)
}

func (c *CharacterController) Update(ctx echo.Context) error {
	// Create new models to hold the updated character
	updatedCharacter := new(models.Character)
	// Bind the new models to the request body
	if err := ctx.Bind(&updatedCharacter); err != nil {
		return returnErrorCode(ctx, http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return returnErrorCode(ctx, http.StatusBadRequest, err)
	}

	// Assign the updated character the current ID
	updatedCharacter.ID = id

	// Update the existing character in the store with the updated information
	err = c.Store.Update(updatedCharacter)
	if err != nil {
		return returnErrorCode(ctx, http.StatusNotFound, err)
	}

	return c.Get(ctx)
}

func (c *CharacterController) LevelUp(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return returnErrorCode(ctx, http.StatusBadRequest, err)
	}

	character, err := c.Store.Get(id)
	if err != nil {
		return returnErrorCode(ctx, http.StatusNotFound, err)
	}

	err = c.Store.LevelUp(character)
	if err != nil {
		return returnErrorCode(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, character)
}

func (c *CharacterController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return returnErrorCode(ctx, http.StatusBadRequest, err)
	}

	err = c.Store.Delete(id)
	if err != nil {
		return returnErrorCode(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, "character successfully deleted")
}

func returnErrorCode(ctx echo.Context, code int, err error) error {
	ctx.Response().Status = code
	return echo.NewHTTPError(code, err.Error())
}
