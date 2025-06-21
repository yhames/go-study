package network

import "github.com/gin-gonic/gin"

func (network *Network) ResponseOk(context *gin.Context, data interface{}) {
	context.JSON(200, data)
}

func (network *Network) ResponseFailed(context *gin.Context, data interface{}) {
	context.JSON(500, data)
}
