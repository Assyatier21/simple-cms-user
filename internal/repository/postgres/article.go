package postgres

import (
	database "cms/database/queries"
	m "cms/models"
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetArticles(ctx echo.Context, limit int, offset int) ([]m.Article, error) {
	var (
		articles []m.Article
		rows     *sql.Rows
		err      error
	)

	rows, err = r.db.Query(database.GetArticles, limit, offset)
	if err != nil {
		log.Println("[GetArticles] can't get list of articles, err:", err.Error())
		return nil, err
	}

	for rows.Next() {
		var temp = m.Article{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.HtmlContent, &temp.CategoryID, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[GetArticles] failed to scan article, err :", err.Error())
			return nil, err
		}
		articles = append(articles, temp)
	}

	if len(articles) > 0 {
		return articles, nil
	} else {
		return []m.Article{}, nil
	}

}
func (r *repository) GetArticleDetails(ctx echo.Context, id int) (m.Article, error) {
	var (
		article m.Article
		err     error
	)

	err = r.db.QueryRow(database.GetArticleDetails, id).Scan(&article.Id, &article.Title, &article.Slug, &article.HtmlContent, &article.CategoryID, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		log.Println("[GetArticleDetails] failed to scan article, err:", err.Error())
		return m.Article{}, err
	}

	return article, nil
}
