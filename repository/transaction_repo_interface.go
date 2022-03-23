package repository

import "WMB/model"

type TransactionRepo interface {
	InsertTransaction(transaction model.Transaction) int
	GetLastId() int
	GetNextId() int
	AllPayment() []model.Transaction
	UpdatePayment(id int)
}
