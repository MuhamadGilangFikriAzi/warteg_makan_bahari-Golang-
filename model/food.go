package model

type Food struct {
	Id       int    `json:"food_id"`
	FoodName string `db:"food_name" json:"food_name"`
	Price    int    `json:"food_price"`
	Code     string `json:"food_code"`
}
