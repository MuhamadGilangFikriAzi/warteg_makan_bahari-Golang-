package repository

import "WMB/model"

type TransactionDetailRepo interface {
	InsertTransactionDetail(transactionId int, dataDetail []model.TransactionDetail)
	GetLastIdDetailId() int
	GetNextDetailId() int
}
