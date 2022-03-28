package middleware

import (
	"WMB/delivery/commonresp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware(requiredToken string) gin.HandlerFunc {
	if requiredToken == "" {
		panic("API_TOKEN not exist")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("api_token")
		if token == "" {
			e := commonresp.NewErrorMessage(http.StatusUnauthorized, "01", "Unautherized")
			c.Abort()
			c.Error(fmt.Errorf("%s", e.ToJson()))
			return
		}

		if token != requiredToken {
			e := commonresp.NewErrorMessage(http.StatusUnauthorized, "02", "Unautherized")
			c.Abort()
			c.Error(fmt.Errorf("%s", e.ToJson()))
			return
		}

		c.Next()
	}
}
