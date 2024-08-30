package controller

import (
	"dungeons-and-dragons/store"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type RaceController struct {
	Store store.RaceStore
}

func (r *RaceController) GetAll(ctx echo.Context) error {
	races, err := r.Store.GetAll()
	if err != nil {
		return returnErrorCode(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, races)
}

func (r *RaceController) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return returnErrorCode(ctx, http.StatusBadRequest, err)
	}

	race, err := r.Store.Get(id)
	if err != nil {
		return returnErrorCode(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, race)
}
