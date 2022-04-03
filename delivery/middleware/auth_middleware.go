package middleware

import (
	"WMB/authenticator"
	"WMB/delivery/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthTokenMiddleware struct {
	acctToken authenticator.Token
}

func NewAuthTokenMiddleware(configToken authenticator.Token) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		acctToken: configToken,
	}
}

func (a *AuthTokenMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			c.Next()
		} else {

			h := authHeader{}
			err := c.ShouldBindHeader(&h)
			if err != nil {
				c.JSON(http.StatusConflict, gin.H{
					"Message": err,
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			fmt.Println(tokenString)
			if tokenString == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"http_code": http.StatusUnauthorized,
					"meessage":  "Unautherized",
					"service":   "Token-auth",
				})
				//e := commonresp.NewErrorMessage(http.StatusUnauthorized, "Token-auth", "01", "Unautherized")
				//c.Abort()
				//c.Error(fmt.Errorf("%s", e.ToJson()))
				//c.JSON("Error")
				return
			}
			token, errToken := a.acctToken.VerifAccessToken(tokenString)
			fmt.Println(errToken)
			if errToken != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "not auth",
				})
				c.Abort()
				return
			}
			if token["iss"] == api.AplicationName {
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"http_code": http.StatusUnauthorized,
					"meessage":  "Unautherized",
					"service":   "Token-auth",
				})
				c.Abort()
				return
			}
			c.Next()
		}
	}
}
