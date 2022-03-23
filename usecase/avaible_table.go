package usecase

import (
	"WMB/model"
	"WMB/repository"
)

type AvaibleTableUseCase interface {
	AllAvaibleTable() []model.Table
	OrderTable(id int)
	BookedTable() []model.Table
}

type avaibleTableUseCase struct {
	repo repository.TableRepo
}

func (f *avaibleTableUseCase) AllAvaibleTable() []model.Table {
	return f.repo.GetAllAvabile()
}

func (f *avaibleTableUseCase) BookedTable() []model.Table {
	return f.repo.BookedTable()
}

func (f *avaibleTableUseCase) OrderTable(id int) {
	f.repo.OrderTable(id)
}

func NewAvaibleTableUseCase(repo repository.TableRepo) AvaibleTableUseCase {
	return &avaibleTableUseCase{
		repo: repo,
	}
}
