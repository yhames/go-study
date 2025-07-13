package service

import (
	"chat-ws/src/repository"
	schema2 "chat-ws/src/types/schema"
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

func (s *Service) InsertChatting(roomName, userName, message string) {
	err := s.repository.InsertChatting(userName, message, roomName)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Service) EnterRoom(roomName string) ([]*schema2.Chat, error) {
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

func (s *Service) FindRoomByName(name string) (*schema2.Room, error) {
	room, err := s.repository.FindRoomByName(name)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, nil // Room not found
	}
	return room, nil
}

func (s *Service) FindRoomAll() ([]*schema2.Room, error) {
	rooms, err := s.repository.FindRoomAll()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
