package service

import "crud-server/repository"

type UserService struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
