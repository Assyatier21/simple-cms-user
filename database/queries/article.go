package database

const (
	GetArticles = "SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id ORDER BY a.id LIMIT $1 OFFSET $2"

	GetArticleDetails = "SELECT a.id, a.title, a.slug, a.html_content, c.id, c.title , c.slug, a.created_at, a.updated_at FROM cms_article a JOIN cms_category c ON a.category_id = c.id WHERE a.id = $1"
)
