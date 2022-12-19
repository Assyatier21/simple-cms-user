package api

import (
	m "cms/models"
	"cms/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetCategoryTree(ctx echo.Context) (err error) {
	var (
		datas []m.Category
	)

	datas, err = h.repository.GetCategoryTree(ctx)
	if err != nil {
		log.Println("[Delivery][GetCategoryTree] can't get list of categories, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to get list of categories")
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	categories := make([]interface{}, len(datas))
	for i, v := range datas {
		categories[i] = v
	}
	res := m.SetResponse(http.StatusOK, "success", categories)
	return ctx.JSON(http.StatusOK, res)
}
func (h *handler) GetCategoryByID(ctx echo.Context) (err error) {
	var (
		id int
	)

	if !utils.IsValidNumeric(ctx.FormValue("id")) {
		res := m.SetError(http.StatusBadRequest, "id must be an integer and can't be empty")
		return ctx.JSON(http.StatusBadRequest, res)
	} else {
		id, _ = strconv.Atoi(ctx.FormValue("id"))
	}

	category, err := h.repository.GetCategoryByID(ctx, id)
	if err != nil {
		log.Println("[Delivery][GetCategoryByID] can't get category details, err:", err.Error())
		res := m.SetError(http.StatusInternalServerError, "failed to get category details")
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	var data []interface{}
	data = append(data, category)

	res := m.SetResponse(http.StatusOK, "success", data)
	return ctx.JSON(http.StatusOK, res)
}