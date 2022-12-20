package postgres

import (
	database "cms/database/queries"
	m "cms/models"
	"cms/utils"
	"context"
	"database/sql"
	"log"
)

func (r *repository) GetArticles(ctx context.Context, limit int, offset int) ([]m.ResArticle, error) {
	var (
		articles []m.ResArticle
		rows     *sql.Rows
		err      error
	)

	rows, err = r.db.Query(database.GetArticles, limit, offset)
	if err != nil {
		log.Println("[GetArticles] can't get list of articles, err:", err.Error())
		return nil, err
	}

	for rows.Next() {
		var temp = m.ResArticle{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.HtmlContent, &temp.ResCategory.Id, &temp.ResCategory.Title, &temp.ResCategory.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[GetArticles] failed to scan article, err :", err.Error())
			return nil, err
		}
		temp.CreatedAt = utils.FormattedTime(temp.CreatedAt)
		temp.UpdatedAt = utils.FormattedTime(temp.UpdatedAt)
		articles = append(articles, temp)
	}

	if len(articles) > 0 {
		return articles, nil
	} else {
		return []m.ResArticle{}, nil
	}
}
func (r *repository) GetArticleDetails(ctx context.Context, id int) (m.ResArticle, error) {
	var (
		article m.ResArticle
		err     error
	)

	err = r.db.QueryRow(database.GetArticleDetails, id).Scan(&article.Id, &article.Title, &article.Slug, &article.HtmlContent, &article.ResCategory.Id, &article.ResCategory.Title, &article.ResCategory.Slug, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		log.Println("[GetArticleDetails] failed to scan article, err:", err.Error())
		return m.ResArticle{}, err
	}
	article.CreatedAt = utils.FormattedTime(article.CreatedAt)
	article.UpdatedAt = utils.FormattedTime(article.UpdatedAt)

	return article, nil
}
