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

type ClassController struct {
	Store stores.ClassStore
}

func (c *ClassController) GetAll(ctx echo.Context) error {
	classes, err := c.Store.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, classes)
}

func (c *ClassController) Get(ctx echo.Context) error {
	class, err := c.Store.Get(ctx.Param("id"))

	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, *class)
}

func (c *ClassController) Update(ctx echo.Context) error {
	existingClass, err := c.Store.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	updatedClassRequest := new(requests.ClassRequest)
	if err := ctx.Bind(&updatedClassRequest); err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}
	if updatedClassRequest == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid request body"))
	}

	if updatedClassRequest.Name == "" {
		updatedClassRequest.Name = existingClass.Name
	}
	if updatedClassRequest.Description == "" {
		updatedClassRequest.Description = existingClass.Description
	}

	updatedClass := &models.Class{
		ID:          existingClass.ID,
		Name:        updatedClassRequest.Name,
		Description: updatedClassRequest.Description,
	}

	err = c.Store.Update(updatedClass)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, updatedClass)
}
