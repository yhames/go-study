package network

import (
	"crud-server/service"
	"crud-server/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	initUserRouter     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router      *Network
	userService *service.UserService
}

func newUserRouter(router *Network, userService *service.UserService) *userRouter {
	initUserRouter.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		userRouterInstance.router.engine.POST("/create", userRouterInstance.create)
		userRouterInstance.router.engine.GET("/", userRouterInstance.get)
		userRouterInstance.router.engine.PATCH("/update", userRouterInstance.update)
		userRouterInstance.router.engine.DELETE("/delete", userRouterInstance.delete)
	})
	return userRouterInstance
}

func (u *userRouter) create(context *gin.Context) {
	fmt.Println("userRouter.create")
	u.router.ResponseOk(context, "User created successfully")
}

func (u *userRouter) get(context *gin.Context) {
	fmt.Println("userRouter.get")
	u.router.ResponseOk(context, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Status:  200,
			Message: "User retrieved successfully",
		},
		User: types.User{
			Name:  "John Doe",
			Email: "johndoe@google.com",
			Age:   30,
		},
	})
}

func (u *userRouter) update(context *gin.Context) {
	fmt.Println("userRouter.update")
	u.router.ResponseOk(context, "User updated successfully")
}

func (u *userRouter) delete(context *gin.Context) {
	fmt.Println("userRouter.delete")
	u.router.ResponseOk(context, "User deleted successfully")
}
