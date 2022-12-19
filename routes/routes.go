package routes

import (
	"cms/internal/delivery/api"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes(handler api.Handler) *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	g := e.Group("/v1")
	g.GET("/articles", handler.GetArticles)
	g.GET("/article", handler.GetArticleDetails)
	g.GET("/categories", handler.GetCategoryTree)
	g.GET("/category", handler.GetCategoryByID)

	return e
}
func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
	}))
}
