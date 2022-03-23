package usecase

import (
	"WMB/model"
	"WMB/repository"
)

type FoodOrederUseCase interface {
	InsertTransaction(transaction model.Transaction) int
}

type foodOrderUseCase struct {
	repo repository.TransactionRepo
}

func (f *foodOrderUseCase) InsertTransaction(transaction model.Transaction) int {
	return f.repo.InsertTransaction(transaction)
}

func NewFoodOrderUseCase(repo repository.TransactionRepo) FoodOrederUseCase {
	return &foodOrderUseCase{
		repo: repo,
	}
}
