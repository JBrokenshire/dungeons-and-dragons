package controllers

import (
	"dnd-api/db/stores"
	res "dnd-api/server/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WeaponController struct {
	Store stores.WeaponStore
}

func (w *WeaponController) GetAll(ctx echo.Context) error {
	weapons, err := w.Store.GetAll()
	if err != nil {
		return res.ErrorResponse(ctx, http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, weapons)
}
