package service

import (
	"chat-ws-control/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		repository: rep,
	}
}
