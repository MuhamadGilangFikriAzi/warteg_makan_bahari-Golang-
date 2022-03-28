package repository

import (
	"WMB/delivery/apprequest"
)

type UsersRepo interface {
	UserLogin(email string, password string) (bool, apprequest.LoginResponse)
	AuthToken(token string) (bool, apprequest.LoginResponse)
}
