package main

import (
	"WMB/delivery/api"
	"WMB/delivery/config"
	"WMB/delivery/middleware"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type wartegServer struct {
	router *gin.Engine
	config *config.Config
}

func (w *wartegServer) initHandlers() {
	w.loginUrl()
	w.router.Use(middleware.TokenAuthMiddleware(w.config.InfraManager.SqlDb()))
	w.v1()
}

func (w *wartegServer) loginUrl() {
	loginApiGroup := w.router.Group("/login")
	api.NewLoginApi(loginApiGroup, w.config.UseCaseManager.LoginUseCase())
}

func (w *wartegServer) v1() {
	foodApiGroup := w.router.Group("/food")
	api.NewFoodApi(foodApiGroup, w.config.UseCaseManager)

	transactionApiGroup := w.router.Group("/transaction")
	api.NewTransactionApi(transactionApiGroup, w.config.UseCaseManager)
}

func (w *wartegServer) Run() {
	w.initHandlers()
	err := w.router.Run(w.config.Get("wmb.api.url"))
	if err != nil {
		panic("Server not starting")
	}
}
func Server() AppServer {
	r := gin.Default()
	c := config.New(".", "config")
	return &wartegServer{
		r,
		c,
	}
}
