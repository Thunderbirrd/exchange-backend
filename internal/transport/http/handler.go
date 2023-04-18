package handler

import (
	_ "github.com/Thunderbirrd/exchange-backend/docs"
	"github.com/Thunderbirrd/exchange-backend/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		requests := api.Group("/requests")
		{
			requests.POST("/", h.createRequest)
			requests.POST("/get-all", h.getRequests)
		}

		exchanges := api.Group("/exchanges")
		{
			exchanges.POST("/", h.createExchange)
			exchanges.POST("/accept", h.acceptExchange)
			exchanges.POST("/decline", h.declineExchange)
			exchanges.POST("/by-id", h.getExchange)
			exchanges.GET("/", h.getUsersExchanges)
		}

		currencies := api.Group("/currencies")
		{
			currencies.GET("/", h.getAllCurrencies)
		}

		airports := api.Group("/airports")
		{
			airports.GET("/", h.getAllAirports)
		}
	}

	return router
}
