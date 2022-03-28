package repository

import (
	"WMB/delivery/apprequest"
)

type TransactionDetailRepo interface {
	InsertTransactionDetail(transactionId int, dataDetail []apprequest.TransactionDetailRequest)
	GetLastIdDetailId() int
	GetNextDetailId() int
}
