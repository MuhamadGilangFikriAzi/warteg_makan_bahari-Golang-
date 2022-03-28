package repository

import (
	"WMB/delivery/apprequest"
	"WMB/model"
)

type TransactionRepo interface {
	InsertTransaction(transaction apprequest.TransactionRequest) int
	GetLastId() int
	GetNextId() int
	AllPayment() []model.Transaction
	UpdatePayment(id int)
	CheckUnpaymentTransaction(id int) bool
}
