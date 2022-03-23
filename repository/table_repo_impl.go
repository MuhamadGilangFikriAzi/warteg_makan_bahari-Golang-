package repository

import (
	"WMB/model"
	"github.com/jmoiron/sqlx"
)

type tableRepoImpl struct {
	tableDb *sqlx.DB
}

func (f *tableRepoImpl) GetAllAvabile() []model.Table {
	var table []model.Table
	f.tableDb.Select(&table, "select * from table_warteg")
	return table
}

func (f *tableRepoImpl) BookedTable() []model.Table {
	var table []model.Table
	f.tableDb.Select(&table, "select * from table_warteg where is_ordered = true")
	return table
}

func (f *tableRepoImpl) OrderTable(id int) {
	f.tableDb.Exec("update table_warteg set is_ordered = true where id = $1", id)
}

func NewTableRepo(tableDb *sqlx.DB) TableRepo {
	tableRepo := tableRepoImpl{
		tableDb: tableDb,
	}

	return &tableRepo
}
