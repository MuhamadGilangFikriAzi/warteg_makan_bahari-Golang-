package usecase

import (
	"WMB/model"
	"WMB/repository"
)

type FoodListUseCase interface {
	AllFood() []model.Food
	GetPrice(id int) int
	Search(search string) model.Food
}

type foodListUseCase struct {
	repo repository.FoodRepo
}

func (f *foodListUseCase) AllFood() []model.Food {
	return f.repo.GetAll()
}

func (f *foodListUseCase) GetPrice(id int) int {
	return f.repo.GetPriceById(id)
}

func (f *foodListUseCase) Search(search string) model.Food {
	return f.repo.GetDataSearch(search)
}

func NewFoodListUseCase(repo repository.FoodRepo) FoodListUseCase {
	return &foodListUseCase{
		repo: repo,
	}
}
