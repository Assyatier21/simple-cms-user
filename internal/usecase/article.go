package usecase

import (
	"cms-user/utils"
	"context"

	"log"
)

func (u *usecase) GetArticles(ctx context.Context, limit int, offset int) ([]interface{}, error) {
	var (
		articles []interface{}
	)

	resData, err := u.repository.GetArticles(ctx, limit, offset)
	if err != nil {
		log.Println("[Usecase][GetArticles] can't get list of articles, err:", err.Error())
		return articles, err
	}

	articles = make([]interface{}, len(resData))
	for i, v := range resData {
		utils.FormatTimeResArticle(&v)
		articles[i] = v
	}

	return articles, nil
}
func (u *usecase) GetArticleDetails(ctx context.Context, id int) ([]interface{}, error) {
	var (
		article []interface{}
	)
	resData, err := u.repository.GetArticleDetails(ctx, id)
	if err != nil {
		log.Println("[Usecase][GetArticleDetails] can't get article details, err:", err.Error())
		return article, err
	}
	utils.FormatTimeResArticle(&resData)

	article = append(article, resData)
	return article, nil
}
