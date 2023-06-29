package routes

import (
	"NIST/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {

	auth_routes := router.Group("api/v1/auth")

	auth_routes.POST("/login", handlers.Login)

}
