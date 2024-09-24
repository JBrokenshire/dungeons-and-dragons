package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RaceController struct {
	Store stores.RaceStore
}

func (r *RaceController) GetAll(ctx echo.Context) error {
	races, err := r.Store.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, races)
}

func (r *RaceController) Get(ctx echo.Context) error {
	race, err := r.Store.Get(ctx.Param("id"))
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, race)
}
