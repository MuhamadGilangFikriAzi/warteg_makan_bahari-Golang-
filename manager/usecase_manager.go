package manager

import "WMB/usecase"

type UseCaseManager interface {
	FoodListUseCase() usecase.FoodListUseCase
	FoodOrderUseCase() usecase.FoodOrederUseCase
	AvaibleTableUseCase() usecase.AvaibleTableUseCase
	FoodOrderDetailUseCase() usecase.FoodOrederDetailUseCase
	PaymentUseCase() usecase.PaymentUseCase
}
type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) FoodListUseCase() usecase.FoodListUseCase {
	return usecase.NewFoodListUseCase(u.repo.FoodRepo())
}

func (u *useCaseManager) FoodOrderUseCase() usecase.FoodOrederUseCase {
	return usecase.NewFoodOrderUseCase(u.repo.TransactionRepo(), u.repo.TransactionDetailRepo())
}

func (u *useCaseManager) AvaibleTableUseCase() usecase.AvaibleTableUseCase {
	return usecase.NewAvaibleTableUseCase(u.repo.TableRepo())
}

func (u *useCaseManager) FoodOrderDetailUseCase() usecase.FoodOrederDetailUseCase {
	return usecase.NewFoodOrderDetailUseCase(u.repo.TransactionDetailRepo())
}

func (u *useCaseManager) PaymentUseCase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(u.repo.TransactionRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
