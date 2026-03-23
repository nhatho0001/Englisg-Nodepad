package router

import (
	"app-notepad/internal/controller"
	"app-notepad/internal/middleware"
	"app-notepad/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, s *services.UserService) {
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	userHander := controller.NewUserHander(s)
	custom_middleware := middleware.NewMiddleware(s)
	api_user := r.Group("/user")
	api_user.POST("/login", userHander.UserLogin)
	api_user.POST("/register", userHander.UserSignUp)
	api_chapter := r.Group("/chapter")
	api_chapter.Use(custom_middleware.NewAuthMiddleware())
	api_chapter.GET("/list-chapter", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
}
