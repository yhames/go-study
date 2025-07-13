package network

import (
	"chat-ws/app/service"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	engine *gin.Engine

	service *service.Service

	port string
	ip   string
}

func NewServer(service *service.Service, port string) *Server {
	s := &Server{
		engine:  gin.New(),
		service: service,
		port:    port,
	}
	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())
	s.engine.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	registerServer(s)

	return s
}

func (s *Server) Start() error {
	log.Println("Starting server...")
	s.setServerInfo()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT)
	go func() {
		<-channel // Wait for interrupt signal
		if err := s.service.ServerSet(s.ip+s.port, false); err != nil {
			log.Println(err)
		}

		type ServerInfoEvent struct {
			Ip     string
			Status bool
		}

		event := ServerInfoEvent{
			Ip:     s.ip + s.port,
			Status: false,
		}
		v, err := json.Marshal(event)
		if err != nil {
			log.Println("Failed to marshal server info event:", err)
		}

		ch := make(chan kafka.Event, 1)
		result, err := s.service.Publish("chat", v, ch)
		if err != nil {
			log.Println("Failed to publish server info event:", err)
		}
		log.Println("Published server info event:", result)

		os.Exit(0) // Exit gracefully
	}()
	return s.engine.Run(s.port)
}

func (s *Server) setServerInfo() {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		panic(err.Error())
	}

	var ip net.IP
	for _, addr := range adders {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip = ipnet.IP
				break
			}
		}
	}
	if ip == nil {
		panic("no ip found")
	}
	s.ip = ip.String()
	err = s.service.ServerSet(s.ip+s.port, true)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Server info set successfully:", s.ip+s.port)
}
