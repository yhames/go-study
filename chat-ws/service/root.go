package service

import "chat-ws/repository"

type Service struct {
	repository *repository.Repository
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		repository: rep,
	}
}
