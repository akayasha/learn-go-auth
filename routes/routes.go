package routes

import (
	"github.com/gin-gonic/gin"
	"learn-go-auth/controllers"
)

// SetupRouter sets up all routes for the application
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Authentication routes
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/verify-email", controllers.VerifyEmail)

	return r
}
