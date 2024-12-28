package handler

import (
	"github.com/Diyarjan/TasksManager/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	tasks := router.Group("/tasks")
	{
		tasks.GET("/", h.getAllTasks)
		tasks.POST("/", h.createTask)
		tasks.PUT("/:id", h.updateTask)
		tasks.DELETE("/:id", h.deleteTask)
	}

	return router
}
