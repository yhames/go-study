package service

import (
	"grpc-server/app/config"
	"grpc-server/app/repository"
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
