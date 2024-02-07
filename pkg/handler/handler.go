package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/morpheus228/infotecs_golang_test/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	wallet := router.Group("/api/v1/wallet")
	{
		wallet.POST("/", h.createWallet)
		wallet.POST("/:walletId/send", h.makeTransaction)
		wallet.GET("/:walletId/history", h.getWalletHistory)
		wallet.GET("/:walletId", h.getWallet)
	}

	return router
}
