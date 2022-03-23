package usecase

import (
	"WMB/model"
	"WMB/repository"
)

type FoodOrederDetailUseCase interface {
	InsertTransactionDetail(transactionId int, detail []model.TransactionDetail)
}

type foodOrderDetailUseCase struct {
	repo repository.TransactionDetailRepo
}

func (f *foodOrderDetailUseCase) InsertTransactionDetail(transactionId int, detail []model.TransactionDetail) {
	f.repo.InsertTransactionDetail(transactionId, detail)
}

func NewFoodOrderDetailUseCase(repo repository.TransactionDetailRepo) FoodOrederDetailUseCase {
	return &foodOrderDetailUseCase{
		repo: repo,
	}
}
