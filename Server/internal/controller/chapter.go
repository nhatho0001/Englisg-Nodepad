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
	uidAny, exist := c.Get("uid")
	if !exist {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You cannot access this page!"})
		return
	}

	uidStr, ok := uidAny.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid user ID format in context"})
		return
	}

	uidInt, err := strconv.Atoi(uidStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: invalid user ID"})
		return
	}

	chapter_list, err := ch.Chapter.GetListChapter(c, int32(uidInt))
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
