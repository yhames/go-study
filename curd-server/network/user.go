package network

import (
	"crud-server/service"
	"crud-server/types"
	"github.com/gin-gonic/gin"
	"strconv"
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
		userRouterInstance.router.engine.PATCH("/update/:id", userRouterInstance.update)
		userRouterInstance.router.engine.DELETE("/delete/:id", userRouterInstance.delete)
	})
	return userRouterInstance
}

func (userRouter *userRouter) create(context *gin.Context) {
	var request types.CreateUserRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		userRouter.router.ResponseFailed(context, err)
		return
	}
	err := userRouter.userService.Create(request.ToUser())
	if err != nil {
		userRouter.router.ResponseFailed(context, err)
	}
	userRouter.router.ResponseOk(context, "User created successfully")
}

func (userRouter *userRouter) get(context *gin.Context) {
	users := userRouter.userService.FindAll()
	userRouter.router.ResponseOk(context, users)
}

func (userRouter *userRouter) update(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		userRouter.router.ResponseFailed(context, err)
		return
	}
	var request types.UpdateUserRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		userRouter.router.ResponseFailed(context, err)
		return
	}
	err = userRouter.userService.Update(id, &types.User{
		Name:  request.Name,
		Email: request.Email,
		Age:   request.Age,
	})
	if err != nil {
		userRouter.router.ResponseFailed(context, err)
		return
	}
	userRouter.router.ResponseOk(context, "User updated successfully")
}

func (userRouter *userRouter) delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		userRouter.router.ResponseFailed(context, err)
		return
	}
	err = userRouter.userService.Delete(id)
	if err != nil {
		userRouter.router.ResponseFailed(context, err)
		return
	}
	userRouter.router.ResponseOk(context, "User deleted successfully")
}
