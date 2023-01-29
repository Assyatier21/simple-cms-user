package utils

import (
	m "cms-user/models"
	"log"
	"time"
)

func FormatTimeResArticle(article *m.ResArticle) m.ResArticle {
	article.CreatedAt = FormattedTime(article.CreatedAt)
	article.UpdatedAt = FormattedTime(article.UpdatedAt)
	return *article
}
func FormatTimeResCategory(category *m.Category) m.Category {
	category.CreatedAt = FormattedTime(category.CreatedAt)
	category.UpdatedAt = FormattedTime(category.UpdatedAt)
	return *category
}
func FormattedTime(ts string) string {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Println(err)
		return ""
	}

	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}
