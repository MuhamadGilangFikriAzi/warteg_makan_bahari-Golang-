package repository

import "WMB/model"

type FoodRepo interface {
	GetAll() []model.Food
	GetPriceById(id int) int
	GetDataSearch(search string) model.Food
}
