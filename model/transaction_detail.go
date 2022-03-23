package model

type TransactionDetail struct {
	Id            int
	TransactionId int `db:"transaction_id"`
	FoodId        int `db:"food_id"`
	Qty           int
}
