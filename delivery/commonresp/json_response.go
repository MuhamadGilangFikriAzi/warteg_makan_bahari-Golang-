package commonresp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	c *gin.Context
}

func (j *JsonResponse) SendData(message ResponMessage) {
	j.c.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendError(message ErrorMessage) {
	j.c.JSON(message.HttpCode, message)
}

func NewJsonResponse(c *gin.Context) AppHttpResponse {
	return &JsonResponse{
		c,
	}
}
