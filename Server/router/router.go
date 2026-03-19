package router

import (
	"app-notepad/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	api_user := r.Group("/user")
	api_user.POST("/login", controller.UserLogin)
	// api_user.POST("/register")
}
