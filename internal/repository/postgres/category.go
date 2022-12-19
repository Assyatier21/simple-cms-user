package postgres

import (
	database "cms/database/queries"
	m "cms/models"
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
)

func (r *repository) GetCategoryTree(ctx echo.Context) ([]m.Category, error) {
	var (
		categories []m.Category
		rows       *sql.Rows
		err        error
	)

	rows, err = r.db.Query(database.GetCategoryTree)
	if err != nil {
		log.Println("[GetCategoryTree] can't get list of categories, err:", err.Error())
		return nil, err
	}

	for rows.Next() {
		var temp = m.Category{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[GetCategoryTree] failed to scan category, err :", err.Error())
			return nil, err
		}
		categories = append(categories, temp)
	}

	if len(categories) > 0 {
		return categories, nil
	} else {
		return []m.Category{}, nil
	}
}
func (r *repository) GetCategoryByID(ctx echo.Context, id int) (m.Category, error) {
	var (
		category m.Category
		err      error
	)

	err = r.db.QueryRow(database.GetCategoryByID, id).Scan(&category.Id, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.Println("[GetArticleDetails] failed to scan article, err:", err.Error())
		return m.Category{}, err
	}

	return category, nil
}
