package repository

import (
	"crud-server/types"
	"errors"
	"sort"
	"sync/atomic"
)

var seq int64 = 0

type UserRepository struct {
	UserMap map[int64]*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		UserMap: make(map[int64]*types.User),
	}
}

func (userRepository *UserRepository) Create(user *types.User) error {
	user.Id = nextSeq()
	if (userRepository.UserMap[user.Id]) != nil {
		return errors.New("user already exists")
	}
	userRepository.UserMap[user.Id] = user
	return nil
}

func (userRepository *UserRepository) FindAll() []*types.User {
	users := make([]*types.User, 0, len(userRepository.UserMap))
	for _, user := range userRepository.UserMap {
		users = append(users, user)
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})
	return users
}

func (userRepository *UserRepository) FindById(id int64) (*types.User, error) {
	user, exists := userRepository.UserMap[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (userRepository *UserRepository) Update(id int64, updateUser *types.User) error {
	user, exists := userRepository.UserMap[id]
	if !exists {
		return errors.New("user not found")
	}
	user.Name = updateUser.Name
	user.Email = updateUser.Email
	user.Age = updateUser.Age
	return nil
}

func (userRepository *UserRepository) Delete(id int64) error {
	_, exists := userRepository.UserMap[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(userRepository.UserMap, id)
	return nil
}

func nextSeq() int64 {
	return atomic.AddInt64(&seq, 1)
}
