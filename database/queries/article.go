package database

const (
	GetArticles       = "SELECT * FROM cms_article ORDER BY id LIMIT $1 OFFSET $2"
	GetArticleDetails = "SELECT a.*, c.title AS category_title , c.slug AS category_slug FROM articles a INNER JOIN categories c ON a.category_id = c.id WHERE a.id = $1"
)
