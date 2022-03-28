package manager

import "WMB/repository"

type RepoManager interface {
	FoodRepo() repository.FoodRepo
	TableRepo() repository.TableRepo
	TransactionRepo() repository.TransactionRepo
	TransactionDetailRepo() repository.TransactionDetailRepo
	UsersRepo() repository.UsersRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) FoodRepo() repository.FoodRepo {
	return repository.NewFoodRepo(r.infra.SqlDb())
}

func (r *repoManager) TableRepo() repository.TableRepo {
	return repository.NewTableRepo(r.infra.SqlDb())
}

func (r *repoManager) TransactionRepo() repository.TransactionRepo {
	return repository.NewTransactionRepo(r.infra.SqlDb())
}

func (r *repoManager) TransactionDetailRepo() repository.TransactionDetailRepo {
	return repository.NewTransactionDetailRepo(r.infra.SqlDb())
}

func (r *repoManager) UsersRepo() repository.UsersRepo {
	return repository.NewUsersRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
