package service

import (
	"crud-server/repository"
	"crud-server/types"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userRepository *UserService) Create(user *types.User) error {
	return userRepository.userRepository.Create(user)
}

func (userRepository *UserService) FindAll() []*types.User {
	return userRepository.userRepository.FindAll()
}

func (userRepository *UserService) FindById(id int64) *types.User {
	return userRepository.userRepository.FindById(id)
}

func (userRepository *UserService) Update(id int64, updateUser *types.User) error {
	return userRepository.userRepository.Update(id, updateUser)
}

func (userRepository *UserService) Delete(id int64) error {
	return userRepository.userRepository.Delete(id)
}
