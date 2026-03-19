package router

import (
	"app-notepad/internal/controller"
	"app-notepad/internal/services"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, s *services.UserService) {
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	userHander := controller.NewUserHander(s)
	api_user := r.Group("/user")
	api_user.POST("/login", userHander.UserLogin)
}
