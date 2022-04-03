package api

import (
	"WMB/authenticator"
	"WMB/delivery/apprequest"
	"WMB/usecase"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var AplicationName = "Warung Makan Bahari"
var JwsSignatureMethod = jwt.SigningMethodHS256
var JwsSignatureKey = []byte("P4ssw0rd")

type LoginApi struct {
	usecase   usecase.LoginUsecase
	authToken authenticator.Token
}

func (l *LoginApi) LoginUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest apprequest.LoginRequest
		err := c.ShouldBindJSON(&loginRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Message": "Can't binding json",
			})
			return
		}
		//emailUser := c.Param("email")
		//password := c.Param("password")
		condition, data := l.usecase.Login(loginRequest.Email, loginRequest.Password)
		if !condition {
			c.AbortWithStatusJSON(http.StatusOK, data)
			return
		}
		token, err := l.authToken.CreateToken(data)
		fmt.Println(token)
		fmt.Println(err)
		data.Token = token
		if err != nil {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"message": err,
				"data":    data,
			})
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signid Method invalid")
		} else if method != JwsSignatureMethod {
			return nil, fmt.Errorf("signid Method invalid")
		}

		return JwsSignatureKey, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func NewLoginApi(apiRoute *gin.RouterGroup, usecase usecase.LoginUsecase, configToken authenticator.Token) {
	api := LoginApi{usecase: usecase, authToken: configToken}

	apiRoute.POST("", api.LoginUsers())
}
