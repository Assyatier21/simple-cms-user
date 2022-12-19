package postgres

import (
	m "cms/models"
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	GetArticles(ctx echo.Context, limit int, offset int) ([]m.Article, error)
	GetArticleDetails(ctx echo.Context, id int) (m.Article, error)
	GetCategoryTree(ctx echo.Context) ([]m.Category, error)
	GetCategoryByID(ctx echo.Context, id int) (m.Category, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
