package router

import (
	"app-notepad/internal/controller"
	"app-notepad/internal/middleware"
	"app-notepad/internal/services"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, s *services.UserService, ch *services.ChapterService, vs *services.VocabularyService) {
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	userHander := controller.NewUserHander(s)
	chapterHander := controller.NewChapterHander(ch)
	vocavularyHander := controller.NewVocabularyHandler(vs)

	custom_middleware := middleware.NewMiddleware(s)
	api_user := r.Group("/user")
	api_user.POST("/login", userHander.UserLogin)
	api_user.POST("/register", userHander.UserSignUp)
	api_chapter := r.Group("/chapter")
	api_chapter.Use(custom_middleware.NewAuthMiddleware())

	api_chapter = r.Group("/chapter", custom_middleware.NewAuthMiddleware())
	api_chapter.GET("/list-chapter", chapterHander.GetListChapter)
	api_chapter.GET("/list-vocabulary", chapterHander.GetDetailChapter)
	api_chapter.PUT("/update-chapter", chapterHander.UpdateChapter)
	api_chapter.POST("/create", chapterHander.CreateChapter)

	api_setting := r.Group("/user-setting", custom_middleware.NewAuthMiddleware())
	api_setting.POST(
		"/refresh-token", userHander.UpdateRefreshToken,
	)

	api_vocabulary := r.Group("/vocabulary", custom_middleware.NewAuthMiddleware())
	api_vocabulary.POST("/create", vocavularyHander.CreateVocabulary)
	api_vocabulary.GET("/chapter", vocavularyHander.GetVocabularyOfChapter)
	api_vocabulary.GET("/user", vocavularyHander.GetVocabularyOfUser)
}
