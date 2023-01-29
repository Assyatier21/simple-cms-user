package api

import (
	m "cms-user/models"
	msg "cms-user/models/lib"
	"cms-user/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetCategoryTree(ctx echo.Context) (err error) {
	categories, err := h.usecase.GetCategoryTree(ctx.Request().Context())
	if err != nil {
		log.Println("[Delivery][GetCategoryTree] can't get list of categories, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "list of categories returned successfully", categories)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) GetCategoryDetails(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	category, err := h.usecase.GetCategoryDetails(ctx.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no category found", []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		}
		log.Println("[Delivery][GetCategoryDetails] can't get category details, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "category details returned successfully", category)
	return ctx.JSON(http.StatusOK, res)
}
