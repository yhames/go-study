package service

import (
	"chat-ws/app/repository"
	"chat-ws/app/types/schema"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type Service struct {
	repository *repository.Repository
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		repository: rep,
	}
}

func (s *Service) ServerSet(ip string, available bool) error {
	err := s.repository.ServerSet(ip, available)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s *Service) Publish(topic string, value []byte, ch chan kafka.Event) (kafka.Event, error) {
	event, err := s.repository.Kafka.Publish(topic, value, ch)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return event, nil
}

func (s *Service) PublishServerStatusEvent(ip string, status bool) {
	type ServerInfoEvent struct {
		Ip     string
		Status bool
	}

	event := ServerInfoEvent{
		Ip:     ip,
		Status: status,
	}
	v, err := json.Marshal(event)
	if err != nil {
		log.Println("Failed to marshal server info event:", err)
	}

	ch := make(chan kafka.Event, 1)
	result, err := s.Publish("chat", v, ch)
	if err != nil {
		log.Println("Failed to publish server info event:", err)
	}
	log.Println("Published server info event:", result)
}

func (s *Service) InsertChatting(roomName, userName, message string) {
	err := s.repository.InsertChatting(userName, message, roomName)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Service) EnterRoom(roomName string) ([]*schema.Chat, error) {
	room, err := s.repository.FindRoomByName(roomName)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, nil // Room not found
	}
	chats, err := s.repository.FindChatByRoomName(roomName)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (s *Service) CreateRoom(name string) error {
	return s.repository.CreateRoom(name)
}

func (s *Service) FindRoomByName(name string) (*schema.Room, error) {
	room, err := s.repository.FindRoomByName(name)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, nil // Room not found
	}
	return room, nil
}

func (s *Service) FindRoomAll() ([]*schema.Room, error) {
	rooms, err := s.repository.FindRoomAll()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
