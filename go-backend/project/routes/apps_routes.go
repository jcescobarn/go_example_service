package routes

import (
	"NIST/handlers"
	middleware "NIST/middlewares"

	"github.com/gin-gonic/gin"
)

func AppsRouter(router *gin.Engine) {

	app_routes := router.Group("api/v1/apps")

	app_routes.POST("", middleware.AuthMiddleware(), handlers.AppCreate)
	app_routes.GET("", middleware.AuthMiddleware(), handlers.AppGet)
	app_routes.GET("/:id", middleware.AuthMiddleware(), handlers.AppGetById)
	app_routes.PUT("/:id", middleware.AuthMiddleware(), handlers.AppUpdate)
	app_routes.DELETE("/:id", middleware.AuthMiddleware(), handlers.AppDelete)
}
