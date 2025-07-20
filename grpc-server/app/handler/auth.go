package handler

import (
	"github.com/gin-gonic/gin"
	"grpc-server/app/types"
	"net/http"
)

func (r *Router) Login(c *gin.Context) {
	var request types.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	response, err := r.service.CreateAuth(request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create auth"})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (r *Router) Verify(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}
