package service

import (
	"crud-server/repository"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository  *repository.Repository
	UserService *UserService
}

func NewService(repository *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository:  repository,
			UserService: newUserService(repository.UserRepository),
		}
	})
	return serviceInstance
}
