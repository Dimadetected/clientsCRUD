package handler

import (
	"github.com/Dimadetected/clientsCRUD/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{services: srv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	clients := router.Group("/clients", h.UserIdentity)
	{
		clients.GET("/", h.getAllClients)
		clients.GET("/:id", h.getByIdClient)
		clients.POST("/", h.createClient)
		clients.PUT("/:id", h.updateClient)
		clients.DELETE("/:id", h.deleteClient)
		sales := router.Group("/sales")
		{
			sales.GET("/", h.getAllSales)
			sales.GET("/:id", h.getByIdSale)
		}
	}

	sales := router.Group("/sales", h.UserIdentity)
	{
		sales.POST("/", h.createSale)
		sales.PUT("/:id", h.updateSale)
		sales.DELETE("/:id", h.deleteSale)
	}

	return router
}
