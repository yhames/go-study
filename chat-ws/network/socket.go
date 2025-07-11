package network

import (
	"chat-ws/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  types.SocketBufferSize,
	WriteBufferSize: types.MessageBufferSize,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Room struct {
	Forward chan *message
	Join    chan *client
	Leave   chan *client
	Clients map[*client]bool
}

type message struct {
	Name    string
	Message string
	Time    int64
}

type client struct {
	Send   chan *message
	Room   *Room
	Name   string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan *message),
		Join:    make(chan *client),
		Leave:   make(chan *client),
		Clients: make(map[*client]bool),
	}
}

func (c *client) Read() {
	defer func() {
		if err := c.Socket.Close(); err != nil {
			log.Printf("socket close error: %v", err)
		}
	}()
	for {
		var msg *message
		if err := c.Socket.ReadJSON(&msg); err != nil {
			log.Printf("read error: %v", err)
			return
		}
		msg.Time = time.Now().Unix()
		msg.Name = c.Name
		c.Room.Forward <- msg
	}
}

func (c *client) Write() {
	defer func() {
		if err := c.Socket.Close(); err != nil {
			log.Printf("socket close error: %v", err)
		}
	}()
	for msg := range c.Send {
		err := c.Socket.WriteJSON(msg)
		if err != nil {
			panic(err)
		}
	}
}

func (r *Room) RunInit() {
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
		case client := <-r.Leave:
			if _, ok := r.Clients[client]; ok {
				close(client.Send)
				delete(r.Clients, client)
			}
		case msg := <-r.Forward:
			for client := range r.Clients {
				client.Send <- msg
			}
		}
	}
}

func (r *Room) SocketServe(c *gin.Context) {
	// HTTP 요청을 WebSocket으로 업그레이드
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to upgrade connection: %v", err)
		return
	}

	// 사용자 인증을 위해 쿠키에서 "auth" 값을 가져옴
	cookie, err := c.Request.Cookie("auth")
	if err != nil {
		c.String(http.StatusUnauthorized, "Authentication required")
		return
	}

	// 클라이언트 생성
	client := &client{
		Socket: ws,
		Send:   make(chan *message, types.MessageBufferSize),
		Room:   r,
		Name:   cookie.Value,
	}

	r.Join <- client
	defer func() { r.Leave <- client }()

	go client.Write()
	client.Read()
}
