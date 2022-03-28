package api

import (
	"WMB/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginApi struct {
	usecase usecase.LoginUsecase
}

func (l *LoginApi) LoginUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailUser := c.Query("email")
		password := c.Query("password")
		condition, data := l.usecase.Login(emailUser, password)
		fmt.Printf("%s %s", condition, data)
		c.JSON(http.StatusOK, data)
	}
}

func NewLoginApi(apiRoute *gin.RouterGroup, usecase usecase.LoginUsecase) {
	api := LoginApi{usecase: usecase}

	apiRoute.GET("", api.LoginUsers())
}
