package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ItemController struct {
	Store stores.ItemStore
}

func (c *ItemController) GetAll(ctx echo.Context) error {
	items, err := c.Store.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, items)
}
