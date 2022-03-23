package model

type Transaction struct {
	Id           int
	TableId      int    `db:"table_id"`
	CustomerName string `db:"customer_name"`
	Total        int
	Status       string
}
