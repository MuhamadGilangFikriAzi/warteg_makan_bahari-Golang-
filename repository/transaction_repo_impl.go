package repository

import (
	"WMB/delivery/apprequest"
	"WMB/delivery/utility"
	"WMB/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type transactionRepoImpl struct {
	transactionDb *sqlx.DB
}

func (f *transactionRepoImpl) InsertTransaction(transaction apprequest.TransactionRequest) int {
	//timestamp := time.Now() // format waktu saat ini
	//fmt.Println(timestamp)
	//fmt.Println(transaction.CustomerName)
	tx := f.transactionDb.MustBegin()
	idTransaction := utility.GetUuid()
	var transactionDetailId string
	_, err := tx.Exec("insert into transaction_warteg(id, table_id, customer_name, status) values($1, $2, $3, $4)", idTransaction, transaction.TableId, transaction.CustomerName, "Proses")
	if err != nil {
		fmt.Println(err)
	}
	for _, detail := range transaction.TransactionDetail {
		transactionDetailId = utility.GetUuid()
		_, err = tx.Exec("insert into transaction_detail(id, transaction_id, food_id, qty) values($1, $2, $3, $4)", transactionDetailId, idTransaction, detail.FoodId, detail.Qty)
		if err != nil {
			fmt.Println(err)
		}
	}
	tx.Commit()
	return f.GetLastId()
}

func (f *transactionRepoImpl) GetLastId() int {
	var lastId int
	f.transactionDb.Get(&lastId, "select id from transaction_warteg order by id desc limit 1")
	return lastId
}

func (f *transactionRepoImpl) CheckUnpaymentTransaction(id int) bool {
	var status string
	isAvailable := false
	f.transactionDb.Get(&status, "select status from transaction_warteg where id = $1", id)
	if status == "Finish" {
		isAvailable = true
	}
	return isAvailable
}

func (f *transactionRepoImpl) AllPayment() []model.Transaction {
	var table []model.Transaction
	f.transactionDb.Select(&table, "select * from transaction_warteg where status = $1", "Proses")
	return table
}

func (f *transactionRepoImpl) UpdatePayment(id int) {
	f.transactionDb.Exec("update transaction_warteg set status = $1 where id = $2", "Finish", id)
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
