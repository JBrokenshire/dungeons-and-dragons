package controller

import (
	"dungeons-and-dragons/db/models"
	"dungeons-and-dragons/db/stores"
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
	updatedClass := new(models.Class)
	if err := ctx.Bind(&updatedClass); err != nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, err)
	}
	if updatedClass == nil {
		return res.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid request body"))
	}

	existingClass, err := c.Store.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	updatedClass.ID = existingClass.ID
	if updatedClass.Name == "" {
		updatedClass.Name = existingClass.Name
	}
	if updatedClass.Description == "" {
		updatedClass.Description = existingClass.Description
	}

	err = c.Store.Update(updatedClass)
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, updatedClass)
}
