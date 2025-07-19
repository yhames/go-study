package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type tower struct {
	server *Network
}

var once sync.Once // Ensure that the tower is only registered once

func registerTowerApi(server *Network) {
	once.Do(func() {
		t := &tower{server: server}
		t.server.engine.GET("/server-list", t.serverList)
	})
}

func (t *tower) serverList(c *gin.Context) {
	response(c, http.StatusOK, t.server.service.GetAvailableServerList())
}
