package postgres

import (
	m "cms/models"
	"cms/utils"
)

func FormatTimeResArticle(article *m.ResArticle) m.ResArticle {
	article.CreatedAt = utils.FormattedTime(article.CreatedAt)
	article.UpdatedAt = utils.FormattedTime(article.UpdatedAt)
	return *article
}
func FormatTimeResCategory(Category *m.Category) m.Category {
	Category.CreatedAt = utils.FormattedTime(Category.CreatedAt)
	Category.UpdatedAt = utils.FormattedTime(Category.UpdatedAt)
	return *Category
}
