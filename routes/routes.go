package routes

import (
	"copy_users_for_moodle/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Gin router oluştur
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowMethods = []string{"OPTIONS", "GET", "POST"}
	router.Use(cors.New(config))

	router.GET("/", controllers.TestGo)
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/getUsers", controllers.CopyUsers)

	return router
}
