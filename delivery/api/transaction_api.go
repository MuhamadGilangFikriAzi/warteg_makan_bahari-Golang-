package api

import (
	"WMB/delivery/apprequest"
	"WMB/delivery/commonresp"
	"WMB/manager"
	"fmt"
	"github.com/gin-gonic/gin"
)

type TransactionApi struct {
	usecase manager.UseCaseManager
}

func (t *TransactionApi) StoreTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var transactionReq apprequest.TransactionRequest
		if err := c.ShouldBindJSON(&transactionReq); err != nil {
			fmt.Println(transactionReq)
			commonresp.NewJsonResponse(c).SendData(commonresp.NewResponMessage("200", "All Data Food", transactionReq))
		}
	}
}

func NewTransactionApi(transactionRoute *gin.RouterGroup, usecase manager.UseCaseManager) {
	api := TransactionApi{
		usecase: usecase,
	}
	transactionRoute.POST("/store", api.StoreTransaction())
}
