package usecase

import (
	"WMB/delivery/apprequest"
	"WMB/repository"
)

type FoodOrederUseCase interface {
	InsertTransaction(transaction apprequest.TransactionRequest)
}

type foodOrderUseCase struct {
	repo       repository.TransactionRepo
	repoDetail repository.TransactionDetailRepo
}

func (f *foodOrderUseCase) InsertTransaction(transaction apprequest.TransactionRequest) {
	//transactionId := f.repo.InsertTransaction(transaction)
	//fmt.Println(transaction.CustomerName)
	f.repo.InsertTransaction(transaction)
	//f.repoDetail.InsertTransactionDetail(transactionId, transaction.TransactionDetail)
}

func NewFoodOrderUseCase(repo repository.TransactionRepo, repoDetail repository.TransactionDetailRepo) FoodOrederUseCase {
	return &foodOrderUseCase{
		repo:       repo,
		repoDetail: repoDetail,
	}
}
