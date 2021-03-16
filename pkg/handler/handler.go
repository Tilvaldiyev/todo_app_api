package handler

import (
	"github.com/gin-gonic/gin"
	"todo_app_api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}



func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := api.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllLists)
				items.GET("/:item_id", h.getItem)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:items_id", h.deleteItem)
			}
		}
	}
	return router
}