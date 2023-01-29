package routes

import (
	"cms-user/internal/delivery/api"
	"cms-user/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes(handler api.DeliveryHandler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	g := e.Group("/v1")
	g.GET(utils.PATH_ARTICLES, handler.GetArticles)
	g.GET(utils.PATH_ARTICLE, handler.GetArticleDetails)
	g.GET(utils.PATH_CATEGORIES, handler.GetCategoryTree)
	g.GET(utils.PATH_CATEGORY, handler.GetCategoryDetails)

	return e
}
func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))
}
