package repository

import (
	"WMB/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type foodRepoImpl struct {
	foodDb *sqlx.DB
}

func (f *foodRepoImpl) GetAll() []model.Food {
	var food []model.Food
	f.foodDb.Select(&food, "select * from food")
	for _, val := range food {
		fmt.Println(val)
	}
	return food
}

func (f *foodRepoImpl) GetPriceById(id int) int {
	var price int
	f.foodDb.Get(&price, "select price from food where id = $1", id)
	return price
}

func (f *foodRepoImpl) GetDataSearch(search string) model.Food {
	var food model.Food
	fmt.Println(search)
	query := `select * from food where food_name = $1 or code = $1 limit 1`
	f.foodDb.Get(&food, query, search)
	return food
}

func NewFoodRepo(foodDb *sqlx.DB) FoodRepo {
	foodRepo := foodRepoImpl{
		foodDb: foodDb,
	}

	return &foodRepo
}
