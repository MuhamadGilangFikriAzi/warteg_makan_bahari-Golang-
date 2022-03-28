package apprequest

type TransactionRequest struct {
	CustomerName      string                     `json:"customer_name"`
	TableId           int                        `json:"table_id"`
	TransactionDetail []TransactionDetailRequest `json:"transaction_detail"`
}

type TransactionDetailRequest struct {
	FoodId int `json:"food_id"`
	Qty    int `json:"qty"`
}
