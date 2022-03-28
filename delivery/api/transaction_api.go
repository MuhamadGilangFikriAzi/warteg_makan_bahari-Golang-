package api

import (
	"WMB/delivery/apprequest"
	"WMB/manager"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionApi struct {
	usecase manager.UseCaseManager
}

func (t *TransactionApi) StoreTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var transactionReq apprequest.TransactionRequest
		if err := c.ShouldBindJSON(&transactionReq); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"meessage": "Can not bind json",
			})
			return
			//commonresp.NewJsonResponse(c).SendData(commonresp.NewResponMessage("200", "Insert Succesfull", transactionReq))
		}
		t.usecase.FoodOrderUseCase().InsertTransaction(transactionReq)
		c.JSON(http.StatusOK, transactionReq)
		return
	}
}

func NewTransactionApi(transactionRoute *gin.RouterGroup, usecase manager.UseCaseManager) {
	api := TransactionApi{
		usecase: usecase,
	}
	transactionRoute.POST("/store", api.StoreTransaction())
}
