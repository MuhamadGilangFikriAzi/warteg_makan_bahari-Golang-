package repository

import (
	"WMB/model"
	"github.com/jmoiron/sqlx"
)

type transactionDetailRepoImpl struct {
	transactionDetailDb *sqlx.DB
}

func (f *transactionDetailRepoImpl) InsertTransactionDetail(transactionId int, dataDetail []model.TransactionDetail) {

	for _, detail := range dataDetail {
		f.transactionDetailDb.Exec("insert into transaction_detail(id, transaction_id, food_id, qty) values($1, $2, $3, $4)", f.GetNextDetailId(), transactionId, detail.FoodId, detail.Qty)
	}
	//f.transactionDetailDb.Exec("insert into transaction_warteg(id, table_id, customer_name, total, status) values($1, $2, $3, $4, $5)", )
}

func (f *transactionDetailRepoImpl) GetLastIdDetailId() int {
	var lastId int
	f.transactionDetailDb.Get(&lastId, "select id from transaction_detail order by id desc limit 1")
	return lastId
}

func (f *transactionDetailRepoImpl) GetNextDetailId() int {
	return f.GetLastIdDetailId() + 1
}

func NewTransactionDetailRepo(tableDb *sqlx.DB) TransactionDetailRepo {
	tableRepo := transactionDetailRepoImpl{
		transactionDetailDb: tableDb,
	}

	return &tableRepo
}
