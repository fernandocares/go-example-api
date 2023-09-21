package router

import (
	"github.com/gin-gonic/gin"
	"go-example-api/config"
)

func Init(init config.Initialization, router *gin.Engine) *gin.Engine {
	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	}
	return router
}
