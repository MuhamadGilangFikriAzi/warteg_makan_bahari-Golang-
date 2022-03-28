package apprequest

type TransactionRequest struct {
	CustomerName      string                     `json:"customer_name" binding="required"`
	TableId           int                        `json:"table_id" binding="required"`
	TransactionDetail []TransactionDetailRequest `json:"transaction_detail"`
}

type TransactionDetailRequest struct {
	FoodId int `json:"food_id" binding="required"`
	Qty    int `json:"qty" binding="required"`
}
