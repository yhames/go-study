package service

import (
	"grpc-server/app/config"
	"grpc-server/app/repository"
	"log"
)

type Service struct {
	config     *config.Config
	repository *repository.Repository
}

func NewService(config *config.Config, repository *repository.Repository) (*Service, error) {
	return &Service{
		config:     config,
		repository: repository,
	}, nil
}

func (s *Service) CreateAuth(name string) (interface{}, error) {
	auth, err := s.repository.CreateAuth(name)
	if err != nil {
		log.Println(err)
	}
	return auth, err
}
