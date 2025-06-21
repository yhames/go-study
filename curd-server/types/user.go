package types

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Request

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}

func (createUserRequest *CreateUserRequest) ToUser() *User {
	return &User{
		Name:  createUserRequest.Name,
		Email: createUserRequest.Email,
		Age:   createUserRequest.Age,
	}
}

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}

// Response

type GetUsersResponse struct {
	Users []*User `json:"users,omitempty"`
}
