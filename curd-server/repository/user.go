package repository

import "crud-server/types"

type UserRepository struct {
	UserMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{}
}
