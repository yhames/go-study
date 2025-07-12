package service

import (
	"chat-ws/repository"
	"chat-ws/types/schema"
)

type Service struct {
	repository *repository.Repository
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		repository: rep,
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
