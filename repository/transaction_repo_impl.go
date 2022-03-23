package repository

import (
	"WMB/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type transactionRepoImpl struct {
	transactionDb *sqlx.DB
}

func (f *transactionRepoImpl) InsertTransaction(transaction model.Transaction) int {
	tx := f.transactionDb.MustBegin()
	_, err := tx.Exec("insert into transaction_warteg(id, table_id, customer_name, total, status) values($1, $2, $3, $4, $5)", f.GetNextId(), transaction.TableId, transaction.CustomerName, transaction.Total, transaction.Status)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
	return f.GetLastId()
}

func (f *transactionRepoImpl) GetLastId() int {
	var lastId int
	f.transactionDb.Get(&lastId, "select id from transaction_warteg order by id desc limit 1")
	return lastId
}

func (f *transactionRepoImpl) AllPayment() []model.Transaction {
	var table []model.Transaction
	f.transactionDb.Select(&table, "select * from transaction_warteg where status = $1", "Proses")
	return table
}

func (f *transactionRepoImpl) UpdatePayment(id int) {
	f.transactionDb.Exec("update transaction_warteg set status = $1 where id = $2", "Finis", id)
	var tableId int
	f.transactionDb.Get(&tableId, "select table_id from transaction_warteg where id = $1", id)
	f.transactionDb.Exec("update table_warteg set is_ordered = false where id = $1", tableId)
}

func (f *transactionRepoImpl) GetNextId() int {
	return f.GetLastId() + 1
}

func NewTransactionRepo(tableDb *sqlx.DB) TransactionRepo {
	tableRepo := transactionRepoImpl{
		transactionDb: tableDb,
	}

	return &tableRepo
}
