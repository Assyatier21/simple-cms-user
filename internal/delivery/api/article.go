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

func (h *handler) GetArticles(ctx echo.Context) (err error) {
	var (
		limit  int
		offset int
	)

	limit = 100
	if ctx.FormValue("limit") != "" {
		limit, err = strconv.Atoi(ctx.FormValue("limit"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_ID)
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	offset = 0
	if ctx.FormValue("offset") != "" {
		offset, err = strconv.Atoi(ctx.FormValue("offset"))
		if err != nil {
			res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_OFFSET)
			return ctx.JSON(http.StatusBadRequest, res)
		}
	}

	articles, err := h.usecase.GetArticles(ctx.Request().Context(), limit, offset)
	if err != nil {
		log.Println("[Delivery][GetArticles] can't get list of articles, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "list of articles returned successfully", articles)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) GetArticleDetails(ctx echo.Context) (err error) {
	var (
		id int
	)

	id, err = strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, msg.ERROR_FORMAT_EMPTY_ID)
		return ctx.JSON(http.StatusBadRequest, res)
	}

	article, err := h.usecase.GetArticleDetails(ctx.Request().Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "no article found", []interface{}{})
			return ctx.JSON(http.StatusOK, res)
		}
		log.Println("[Delivery][GetArticleDetails] can't get article details, err:", err.Error())
		res := m.SetError(http.StatusBadRequest, utils.STATUS_FAILED, err.Error())
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	res := m.SetResponse(http.StatusOK, utils.STATUS_SUCCESS, "article details returned successfully", article)
	return ctx.JSON(http.StatusOK, res)
}
