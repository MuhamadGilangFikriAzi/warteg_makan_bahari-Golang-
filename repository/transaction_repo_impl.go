package repository

import (
	"WMB/delivery/apprequest"
	"WMB/delivery/utility"
	"WMB/logger"
	"WMB/model"
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
	tx.Exec("insert into transaction_warteg(id, table_id, customer_name, status) values($1, $2, $3, $4)", idTransaction, transaction.TableId, transaction.CustomerName, "Proses")
	for _, detail := range transaction.TransactionDetail {
		transactionDetailId = utility.GetUuid()
		tx.Exec("insert into transaction_detail(id, transaction_id, food_id, qty) values($1, $2, $3, $4)", transactionDetailId, idTransaction, detail.FoodId, detail.Qty)
	}
	err := tx.Commit()
	if err != nil {
		logger.Log.Err(err).Msg("Error Insert Transaction")
	}
	return f.GetLastId()
}

func (f *transactionRepoImpl) GetLastId() int {
	var lastId int
	err := f.transactionDb.Get(&lastId, "select id from transaction_warteg order by id desc limit 1")
	if err != nil {
		logger.Log.Err(err).Msg("Error Get Last Id")
	}
	return lastId
}

func (f *transactionRepoImpl) CheckUnpaymentTransaction(id int) bool {
	var status string
	isAvailable := false
	err := f.transactionDb.Get(&status, "select status from transaction_warteg where id = $1", id)
	if err != nil {
		logger.Log.Err(err).Msg("Error Check Unpayment")
	}
	if status == "Finish" {
		isAvailable = true
	}
	return isAvailable
}

func (f *transactionRepoImpl) AllPayment() []model.Transaction {
	var table []model.Transaction
	err := f.transactionDb.Select(&table, "select * from transaction_warteg where status = $1", "Proses")
	if err != nil {
		logger.Log.Err(err).Msg("Error Get All Payment")
	}
	return table
}

func (f *transactionRepoImpl) UpdatePayment(id int) {
	var tableId int
	tx := f.transactionDb.MustBegin()
	tx.Exec("update transaction_warteg set status = $1 where id = $2", "Finish", id)
	tx.Get(&tableId, "select table_id from transaction_warteg where id = $1", id)
	tx.Exec("update table_warteg set is_ordered = false where id = $1", tableId)
	err := tx.Commit()
	if err != nil {
		logger.Log.Err(err).Msg("Error Updare payment")
	}
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
