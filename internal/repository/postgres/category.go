package postgres

import (
	database "cms/database/queries"
	m "cms/models"
	"cms/utils"
	"context"
	"database/sql"
	"log"
)

func (r *repository) GetCategoryTree(ctx context.Context) ([]m.Category, error) {
	var (
		categories []m.Category
		rows       *sql.Rows
		err        error
	)

	rows, err = r.db.Query(database.GetCategoryTree)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrNotFound
		} else {
			log.Println("[GetCategoryTree] can't get list of categories, err:", err.Error())
			return nil, err
		}
	}

	for rows.Next() {
		var temp = m.Category{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Slug, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("[GetCategoryTree] failed to scan category, err :", err.Error())
			return nil, err
		}
		FormatTimeResCategory(&temp)
		categories = append(categories, temp)
	}

	if len(categories) > 0 {
		return categories, nil
	} else {
		return []m.Category{}, nil
	}
}
func (r *repository) GetCategoryByID(ctx context.Context, id int) (m.Category, error) {
	var (
		category m.Category
		err      error
	)

	err = r.db.QueryRow(database.GetCategoryByID, id).Scan(&category.Id, &category.Title, &category.Slug, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return m.Category{}, utils.ErrNotFound
		} else {
			log.Println("[GetCategoryByID] failed to scan category, err:", err.Error())
			return m.Category{}, err
		}
	}
	FormatTimeResCategory(&category)
	return category, nil
}
