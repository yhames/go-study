package network

import (
	types2 "chat-ws/app/types"
	"chat-ws/app/types/schema"
	"github.com/gin-gonic/gin"
	"net/http"
)

type api struct {
	server *Server
}

func registerServer(server *Server) {
	a := &api{server: server}

	server.engine.GET("/room-list", a.getRoomList)
	server.engine.POST("/make-room", a.postMakeRoom)
	server.engine.GET("/room", a.getRoom)
	server.engine.GET("/enter-room", a.getEnterRoom)

	r := NewRoom(server.service)
	go r.RunInit()
	server.engine.GET("/room-chat", r.SocketServe)
}

func (a *api) getRoomList(c *gin.Context) {
	result, err := a.server.service.FindRoomAll()
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		response(c, http.StatusOK, []*schema.Room{})
		return
	}
	response(c, http.StatusOK, result)
}

func (a *api) postMakeRoom(c *gin.Context) {
	var req types2.BodyRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := a.server.service.CreateRoom(req.Name); err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	response(c, http.StatusOK, "Success")
}

func (a *api) getRoom(c *gin.Context) {
	var req types2.FormRoomRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusBadRequest, err.Error())
		return
	}
	room, err := a.server.service.FindRoomByName(req.Name)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	if room == nil {
		response(c, http.StatusNotFound, "Room not found")
		return
	}
	response(c, http.StatusOK, room)
}

func (a *api) getEnterRoom(c *gin.Context) {
	var req types2.FormRoomRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response(c, http.StatusBadRequest, err.Error())
		return
	}
	chats, err := a.server.service.EnterRoom(req.Name)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	if chats == nil {
		response(c, http.StatusNotFound, "No room found with that name")
		return
	}
	response(c, http.StatusOK, chats)
}

func response(c *gin.Context, s int, res interface{}, data ...string) {
	c.JSON(s, types2.NewRes(s, res, data...))
}
