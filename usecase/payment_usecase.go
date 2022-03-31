package usecase

import (
	"WMB/model"
	"WMB/repository"
)

type PaymentUseCase interface {
	AllPayment() []model.Transaction
	UpdatePayment(id int)
	CheckUnpaymentTransaction(id int) bool
}

type paymentUseCase struct {
	repo repository.TransactionRepo
}

func (f *paymentUseCase) AllPayment() []model.Transaction {
	return f.repo.AllPayment()
}

func (f *paymentUseCase) UpdatePayment(id int) {
	f.repo.UpdatePayment(id)
}

func (f *paymentUseCase) CheckUnpaymentTransaction(id int) bool {
	return f.repo.CheckUnpaymentTransaction(id)
}

func NewPaymentUseCase(repo repository.TransactionRepo) PaymentUseCase {
	return &paymentUseCase{
		repo: repo,
	}
}
