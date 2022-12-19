package database

const (
	GetArticles       = "SELECT * FROM cms_article ORDER BY id LIMIT $1 OFFSET $2"
	GetArticleDetails = "SELECT * FROM cms_article WHERE id = $1"
)
