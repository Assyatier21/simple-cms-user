package usecase

import (
	"cms-user/utils"
	"context"
	"log"
)

func (u *usecase) GetCategoryTree(ctx context.Context) ([]interface{}, error) {
	var (
		categories []interface{}
	)

	resData, err := u.repository.GetCategoryTree(ctx)
	if err != nil {
		log.Println("[Usecase][GetCategoryTree] can't get list of categories, err:", err.Error())
		return categories, err
	}

	categories = make([]interface{}, len(resData))
	for i, v := range resData {
		utils.FormatTimeResCategory(&v)
		categories[i] = v
	}

	return categories, nil
}
func (u *usecase) GetCategoryDetails(ctx context.Context, id int) ([]interface{}, error) {
	var (
		category []interface{}
	)

	resData, err := u.repository.GetCategoryDetails(ctx, id)
	if err != nil {
		log.Println("[Usecase][GetCategoryDetails] can't get category details, err:", err.Error())
		return category, err
	}
	utils.FormatTimeResCategory(&resData)

	category = append(category, resData)
	return category, nil
}
