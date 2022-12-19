package database

const (
	GetCategoryTree = "SELECT * FROM cms_category ORDER BY id"
	GetCategoryByID = "SELECT * FROM cms_category WHERE id = $1"
)
