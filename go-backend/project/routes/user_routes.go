package routes

import (
	"NIST/handlers"
	middleware "NIST/middlewares"

	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.Engine) {

	user_routes := router.Group("api/v1/user")

	user_routes.POST("", middleware.AuthMiddleware(), handlers.UserCreate)
	user_routes.GET("", middleware.AuthMiddleware(), handlers.UsersGet)
	user_routes.GET("/:username", middleware.AuthMiddleware(), handlers.UserGetByUsername)
	user_routes.PUT("/:username", middleware.AuthMiddleware(), handlers.UserUpdate)
	user_routes.DELETE("/:username", middleware.AuthMiddleware(), handlers.UserDelete)
}
