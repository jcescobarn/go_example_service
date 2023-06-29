package routes

import (
	"NIST/handlers"
	middleware "NIST/middlewares"

	"github.com/gin-gonic/gin"
)

func AppRulesRouter(router *gin.Engine) {

	app_rules_routes := router.Group("api/v1/app_rules")
	app_rules_routes.GET("", middleware.AuthMiddleware(), handlers.GetAllRules)
	app_rules_routes.GET("/severity/:severity", middleware.AuthMiddleware(), handlers.GetAllRulesBySeverity)
	app_rules_routes.GET("/:app_id", middleware.AuthMiddleware(), handlers.GetFixedAppRules)
	app_rules_routes.GET("/:app_id/:severity", middleware.AuthMiddleware(), handlers.GetAppRulesWithoutFixed)
	app_rules_routes.DELETE("/:app_id/:rule_id", middleware.AuthMiddleware(), handlers.DeleteAppRuleByApp)
	app_rules_routes.POST("", middleware.AuthMiddleware(), handlers.AppRuleCreate)

}
