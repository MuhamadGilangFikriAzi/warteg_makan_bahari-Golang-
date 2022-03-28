package api

import (
	"WMB/delivery/commonresp"
	"WMB/manager"
	"github.com/gin-gonic/gin"
)

type FoodApi struct {
	usecase manager.UseCaseManager
}

func (f *FoodApi) GetAllFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := f.usecase.FoodListUseCase().AllFood()
		commonresp.NewJsonResponse(c).SendData(commonresp.NewResponMessage("200", "All Data Food", data))
	}
}

func (f *FoodApi) SearchFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		searchQuery := c.Query("search_key")
		data := f.usecase.FoodListUseCase().Search(searchQuery)
		commonresp.NewJsonResponse(c).SendData(commonresp.NewResponMessage("200", "Get data food", data))
	}
}

func NewFoodApi(foodrRoute *gin.RouterGroup, usecase manager.UseCaseManager) {
	api := FoodApi{
		usecase: usecase,
	}
	foodrRoute.GET("/getall", api.GetAllFood())
	foodrRoute.GET("/searchfood", api.SearchFood())
}