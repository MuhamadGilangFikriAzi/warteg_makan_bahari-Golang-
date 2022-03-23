package model

type Table struct {
	Id          int
	TableNumber int  `db:"table_number"`
	IsOrdered   bool `db:"is_ordered"`
}
