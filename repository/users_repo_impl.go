package repository

import (
	"WMB/delivery/apprequest"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	usersDb *sqlx.DB
}

func (u *usersRepo) UserLogin(email string, password string) (bool, apprequest.LoginResponse) {
	var dataUser apprequest.LoginResponse
	err := u.usersDb.Get(&dataUser, "select name, email, token  from users where email = $1 and password = $2 limit 1", email, password)
	if err != nil {
		fmt.Println(err)
		return false, apprequest.LoginResponse{}
	}
	return true, dataUser
}

func (u *usersRepo) AuthToken(token string) (bool, apprequest.LoginResponse) {
	var dataUser apprequest.LoginResponse
	err := u.usersDb.Get(&dataUser, "select name, email, token from users where token = $1", token)
	if err != nil {
		return false, apprequest.LoginResponse{}
	}
	return true, dataUser
}

func NewUsersRepo(usersDb *sqlx.DB) UsersRepo {
	return &usersRepo{usersDb}
}
