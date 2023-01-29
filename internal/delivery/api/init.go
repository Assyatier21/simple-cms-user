package api

import (
	"cms-user/internal/usecase"

	"github.com/labstack/echo/v4"
)

type DeliveryHandler interface {
	GetArticles(ctx echo.Context) (err error)
	GetArticleDetails(ctx echo.Context) (err error)

	GetCategoryTree(ctx echo.Context) (err error)
	GetCategoryDetails(ctx echo.Context) (err error)
}

type handler struct {
	usecase usecase.UsecaseHandler
}

func NewHandler(usecase usecase.UsecaseHandler) DeliveryHandler {
	return &handler{
		usecase: usecase,
	}
}
