package repository

import (
	"WMB/delivery/apprequest"
	"WMB/logger"
	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	usersDb *sqlx.DB
}

func (u *usersRepo) UserLogin(email string, password string) (bool, apprequest.LoginResponse) {
	var dataUser apprequest.LoginResponse
	err := u.usersDb.Get(&dataUser, "select name, email, token  from users where email = $1 and password = $2 limit 1", email, password)
	if err != nil {
		logger.Log.Err(err).Msg("Error User Login")
		return false, apprequest.LoginResponse{}
	}
	return true, dataUser
}

func (u *usersRepo) AuthToken(token string) (bool, apprequest.LoginResponse) {
	var dataUser apprequest.LoginResponse
	err := u.usersDb.Get(&dataUser, "select name, email, token from users where token = $1", token)
	if err != nil {
		logger.Log.Err(err).Msg("Service: Auth-Token")
		return false, apprequest.LoginResponse{}
	}
	return true, dataUser
}

func NewUsersRepo(usersDb *sqlx.DB) UsersRepo {
	return &usersRepo{usersDb}
}
