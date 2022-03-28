package apprequest

type FoodRequest struct {
	FoodId   int    `json:"food_id"`
	FoodName string `json:"food_name"`
	Price    int    `json:"price"`
}
