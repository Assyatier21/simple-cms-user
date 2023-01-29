package usecase

import (
	"cms-user/internal/repository/postgres"
	"context"
)

type UsecaseHandler interface {
	GetArticles(ctx context.Context, limit int, offset int) ([]interface{}, error)
	GetArticleDetails(ctx context.Context, id int) ([]interface{}, error)

	GetCategoryTree(ctx context.Context) ([]interface{}, error)
	GetCategoryDetails(ctx context.Context, id int) ([]interface{}, error)
}

type usecase struct {
	repository postgres.RepositoryHandler
}

func NewUsecase(repository postgres.RepositoryHandler) UsecaseHandler {
	return &usecase{
		repository: repository,
	}
}
