package repository

import "WMB/model"

type TableRepo interface {
	GetAllAvabile() []model.Table
	OrderTable(id int)
	BookedTable() []model.Table
}
