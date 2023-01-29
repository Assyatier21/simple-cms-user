package postgres

import (
	database "cms-user/database/queries"
	m "cms-user/models"
	"context"
	"database/sql"
	"encoding/json"
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
		log.Println("[Repository][GetArticles] can't get list of articles, err:", err.Error())
		return nil, err
	}

	for rows.Next() {
		var (
			temp         m.ResArticle
			byteMetadata []byte
		)

		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.HtmlContent, &temp.ResCategory.Id, &temp.ResCategory.Title, &temp.ResCategory.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[Repository][GetArticles] failed to scan article, err :", err.Error())
			return nil, err
		}

		err = r.db.QueryRow(database.GetMetaData, temp.Id).Scan(&byteMetadata)
		if err != nil {
			log.Println("[Repository][GetArticles] failed to scan metadata, err :", err.Error())
			return nil, err
		}

		json.Unmarshal(byteMetadata, &temp.MetaData)
		articles = append(articles, temp)
	}

	if len(articles) == 0 {
		return []m.ResArticle{}, nil
	}

	return articles, nil
}
func (r *repository) GetArticleDetails(ctx context.Context, id int) (m.ResArticle, error) {
	var (
		article      m.ResArticle
		err          error
		byteMetadata []byte
	)

	err = r.db.QueryRow(database.GetArticleDetails, id).Scan(&article.Id, &article.Title, &article.Slug, &article.HtmlContent, &article.ResCategory.Id, &article.ResCategory.Title, &article.ResCategory.Slug, &article.CreatedAt, &article.UpdatedAt)
	if err != nil {
		log.Println("[Repository][GetArticleDetails] failed to scan article, err:", err.Error())
		return m.ResArticle{}, err
	}

	err = r.db.QueryRow(database.GetMetaData, article.Id).Scan(&byteMetadata)
	if err != nil {
		log.Println("[Repository][GetArticleDetails] failed to scan metadata, err :", err.Error())
		return m.ResArticle{}, err
	}
	json.Unmarshal(byteMetadata, &article.MetaData)

	return article, nil
}
