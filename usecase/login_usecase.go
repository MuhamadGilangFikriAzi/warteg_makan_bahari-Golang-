package usecase

import (
	"WMB/delivery/apprequest"
	"WMB/repository"
)

type LoginUsecase interface {
	Login(username string, password string) (bool, apprequest.LoginResponse)
}

type loginUsecase struct {
	repo repository.UsersRepo
}

func (l *loginUsecase) Login(username string, password string) (bool, apprequest.LoginResponse) {
	return l.repo.UserLogin(username, password)
}

func NewLoginUsecase(repo repository.UsersRepo) LoginUsecase {
	return &loginUsecase{
		repo,
	}
}
