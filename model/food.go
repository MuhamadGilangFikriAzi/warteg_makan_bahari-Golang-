package model

type Food struct {
	Id       int
	FoodName string `db:"food_name"`
	Price    int
	Code     string
}
