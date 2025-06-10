package network

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	initUserRouter     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
}

func newUserRouter(router *Network) *userRouter {
	initUserRouter.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		userRouterInstance.router.engine.POST("/create", userRouterInstance.create)
		userRouterInstance.router.engine.GET("/", userRouterInstance.get)
		userRouterInstance.router.engine.PATCH("/update", userRouterInstance.update)
		userRouterInstance.router.engine.DELETE("/delete", userRouterInstance.delete)
	})
	return userRouterInstance
}

func (u *userRouter) create(context *gin.Context) {
}

func (u *userRouter) get(context *gin.Context) {
}

func (u *userRouter) update(context *gin.Context) {
}

func (u *userRouter) delete(context *gin.Context) {
}
