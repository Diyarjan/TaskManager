package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllTasks(c *gin.Context) {
	// skip TODO Process input params

	// TODO Send params to service
	lists, err := h.services.Tasks.ListTasks(c.Request.Context())
	if err != nil {
		log.Printf("can't list tasks. Err: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, lists)
		return
	}
	// TODO Respond to user
	c.JSON(http.StatusOK, lists)
}

func (h *Handler) createTask(c *gin.Context) {

}

func (h *Handler) updateTask(c *gin.Context) {

}

func (h *Handler) deleteTask(c *gin.Context) {

}
