package main

import (
	"NIST/config"
	"NIST/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvFile()
	config.ConnectToDB()
}

func main() {
	r := gin.Default()

	routes.AppsRouter(r)
	routes.UsersRouter(r)
	routes.AppRulesRouter(r)
	routes.AuthRouter(r)

	r.Run()
}
