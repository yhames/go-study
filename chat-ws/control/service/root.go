package service

import (
	"chat-ws-control/repository"
	"chat-ws-control/types/table"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type Service struct {
	repository    *repository.Repository
	AvgServerList map[string]bool
}

func NewService(rep *repository.Repository) *Service {
	s := &Service{
		repository:    rep,
		AvgServerList: make(map[string]bool),
	}
	s.SetServerInfo() // Initialize the server info on creation
	err := s.repository.Kafka.Subscribe("chat")
	if err != nil {
		panic("Failed to subscribe to Kafka topic: " + err.Error())
	}
	go s.loopSubKafka()
	return s
}

func (s *Service) loopSubKafka() {
	for {
		event := s.repository.Kafka.Poll(100)
		switch eventType := event.(type) {
		case *kafka.Message:
			type ServerInfoEvent struct {
				Ip     string `json:"ip"`
				Status bool   `json:"status"`
			}
			var decoder ServerInfoEvent
			if err := json.Unmarshal(eventType.Value, &decoder); err != nil {
				log.Println("Failed to unmarshal Kafka message:", err)
				continue
			}
			log.Printf("ServerInfoEvent: %v", decoder)
			s.AvgServerList[decoder.Ip] = decoder.Status
		case *kafka.Error:
			log.Println("Failed to poll Kafka event:", eventType)
		}
	}
}

func (s *Service) GetAvailableServerList() []string {
	var serverList []string
	for ip, available := range s.AvgServerList {
		if available {
			serverList = append(serverList, ip)
		}
	}
	return serverList
}

func (s *Service) GetAvailableServers() ([]*table.ServerInfo, error) {
	servers, err := s.repository.GetAvailableServers()
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func (s *Service) SetServerInfo() {
	servers, err := s.GetAvailableServers()
	if err != nil {
		panic(err)
	}

	for _, server := range servers {
		s.AvgServerList[server.Ip] = server.Available
	}
}
