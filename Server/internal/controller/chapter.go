package controller

import (
	"app-notepad/internal/services"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChapterHander struct {
	Chapter *services.ChapterService
}

func NewChapterHander(s *services.ChapterService) *ChapterHander {
	return &ChapterHander{
		Chapter: s,
	}
}

func (ch *ChapterHander) GetListChapter(c *gin.Context) {
	uid := c.Param("uid")
	uid_64, err := strconv.ParseInt(uid, 10, 32)
	if err != nil {
		slog.Error("Param is not suitable")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Param is not suitable",
		})
		return
	}

	chapter_list, err := ch.Chapter.GetListChapter(c, int32(uid_64))
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": chapter_list,
	})
}
