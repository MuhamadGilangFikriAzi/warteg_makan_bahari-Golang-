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
	f.InsertManyFood()
	return food
}

func (f *foodRepoImpl) GetPriceById(id int) int {
	var price int
	f.foodDb.Get(&price, "select price from food where id = $1", id)
	return price
}

func (f *foodRepoImpl) InsertManyFood() {
	//var price int
	for i := 1000000; i < 1000000; i++ {
		f.foodDb.Exec("insert into data_dumy_food(id,name, price) values($1, $2, $3)", i+1, fmt.Sprintf("Ayam %d", i), 10000+i)
	}
	//f.foodDb.Get(&price, "select price from food where id = $1", id)
	//return price
}

func (f *foodRepoImpl) GetDataSearch(search string) model.Food {
	var food model.Food
	fmt.Println(search)
	query := `select * from food where upper(food_name) like upper($1) or code like $1 limit 1`
	f.foodDb.Get(&food, query, fmt.Sprintf("%s%%", search)) // penggunaan like
	return food
}

func NewFoodRepo(foodDb *sqlx.DB) FoodRepo {
	foodRepo := foodRepoImpl{
		foodDb: foodDb,
	}

	return &foodRepo
}
