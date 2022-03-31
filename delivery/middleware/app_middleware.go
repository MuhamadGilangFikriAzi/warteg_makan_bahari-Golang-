package middleware

import (
	"WMB/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func AuthUser(requiredToken string, sqlxdb *sqlx.DB) bool {
	var countAuthUser int
	err := sqlxdb.Get(&countAuthUser, "select count(*) from users where token = $1", requiredToken)
	if err != nil {
		logger.Log.Err(err).Msg("Database not conected")
		panic("database not connected")
	}
	if countAuthUser == 0 {
		logger.Log.Err(err).Msg("Token Auth not found")
		return true
	}
	return false
}

func TokenAuthMiddleware(sqlxdb *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("api_token")
		if token == "" {
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

		if AuthUser(token, sqlxdb) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"http_code": http.StatusUnauthorized,
				"meessage":  "Unautherized",
				"service":   "Token-auth",
			})
			//e := commonresp.NewErrorMessage(http.StatusUnauthorized, "Token-auth", "02", "Unautherized")
			//c.Abort()
			//c.Error(fmt.Errorf("%s", e.ToJson()))
			return
		}

		c.Next()
	}
}
