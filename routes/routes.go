package routes

import (
	"auth/controllers"
	"auth/middlewares"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.Use(middlewares.AuthMiddleware())
	r.POST("/logout", controllers.Logout)
	return r
}
