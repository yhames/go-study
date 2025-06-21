package repository

import (
	"crud-server/types"
	"sync/atomic"
)

var seq int64 = 0

type UserRepository struct {
	UserMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		UserMap: []*types.User{},
	}
}

func (userRepository *UserRepository) Create(user *types.User) error {
	user.Id = nextSeq()
	userRepository.UserMap = append(userRepository.UserMap, user)
	return nil
}

func (userRepository *UserRepository) FindAll() []*types.User {
	return userRepository.UserMap
}

func (userRepository *UserRepository) FindById(id int64) *types.User {
	for _, u := range userRepository.UserMap {
		if u.Id == id {
			return u
		}
	}
	return nil
}

func (userRepository *UserRepository) Update(id int64, updateUser *types.User) error {
	for i, u := range userRepository.UserMap {
		if u.Id == id {
			userRepository.UserMap[i].Name = updateUser.Name
			userRepository.UserMap[i].Email = updateUser.Email
			userRepository.UserMap[i].Age = updateUser.Age
			return nil
		}
	}
	return nil
}

func (userRepository *UserRepository) Delete(id int64) error {
	for i, u := range userRepository.UserMap {
		if u.Id == id {
			userRepository.UserMap = append(userRepository.UserMap[:i], userRepository.UserMap[i+1:]...)
			return nil
		}
	}
	return nil
}

func nextSeq() int64 {
	return atomic.AddInt64(&seq, 1)
}
