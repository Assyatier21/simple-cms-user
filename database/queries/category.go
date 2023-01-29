package database

const (
	GetCategoryTree = `SELECT * FROM cms_category 
							ORDER BY id`

	GetCategoryDetails = `SELECT * FROM cms_category 
							WHERE id = $1`
)
