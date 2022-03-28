package usecase

import (
	"WMB/delivery/apprequest"
	"WMB/repository"
)

type FoodOrederDetailUseCase interface {
	InsertTransactionDetail(transactionId int, detail []apprequest.TransactionDetailRequest)
}

type foodOrderDetailUseCase struct {
	repo repository.TransactionDetailRepo
}

func (f *foodOrderDetailUseCase) InsertTransactionDetail(transactionId int, detail []apprequest.TransactionDetailRequest) {
	f.repo.InsertTransactionDetail(transactionId, detail)
}

func NewFoodOrderDetailUseCase(repo repository.TransactionDetailRepo) FoodOrederDetailUseCase {
	return &foodOrderDetailUseCase{
		repo: repo,
	}
}
