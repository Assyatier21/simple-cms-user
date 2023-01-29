package postgres

import (
	m "cms-user/models"
	"context"
	"database/sql"
)

type RepositoryHandler interface {
	GetArticles(ctx context.Context, limit int, offset int) ([]m.ResArticle, error)
	GetArticleDetails(ctx context.Context, id int) (m.ResArticle, error)

	GetCategoryTree(ctx context.Context) ([]m.Category, error)
	GetCategoryDetails(ctx context.Context, id int) (m.Category, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryHandler {
	return &repository{
		db: db,
	}
}
