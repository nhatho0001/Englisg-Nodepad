package controller

import (
	"app-notepad/internal/services"
	"app-notepad/internal/store"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type ChapterHander struct {
	Chapter *services.ChapterService
}

type ChapterRequest struct {
	Chapter         store.CreateChapterParams
	List_Vocabulary []store.CreateVocabularyParams
}

type ChapterResponse struct {
	Chapter         store.Chapter
	List_Vocabulary []store.Vocabulary
}

func (c *ChapterRequest) ValidateDataInput() bool {
	if !c.Chapter.Title.Valid || len(c.List_Vocabulary) == 0 {
		return false
	}
	return true
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

func (ch *ChapterHander) GetDetailChapter(c *gin.Context) {
	var query_data SearchFilters
	if err := c.BindQuery(&query_data); err != nil {
		slog.Error("Parse data post is faild!")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Data post server is not parse!",
		})
		return
	}

	chapter, err := ch.Chapter.GetChapterById(c, query_data.ID)
	if err != nil {
		slog.Error("Get Chapter is faild!")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Get Chapter is faild!",
		})
		return
	}

	list_vocabulary, err := ch.Chapter.GetVocabularyOfChapter(c, pgtype.Int4{
		Int32: chapter.ID,
		Valid: true,
	})
	if err != nil {
		slog.Error("Get Vocabulary of Chapter is faild")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Get Vocabulary of Chapter is faild",
		})
		return
	}

	c.JSON(http.StatusOK, ChapterResponse{
		Chapter:         *chapter,
		List_Vocabulary: list_vocabulary,
	})
}

func (ch *ChapterHander) CreateChapter(c *gin.Context) {
	var data ChapterRequest
	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !data.ValidateDataInput() {
		slog.Error("Field request is not have")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Field request is not have",
		})
		return
	}

	uid, exist := c.Get("uid")
	if !exist {
		slog.Error("Error account of you is not verify")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Error account of you is not verify",
		})
		return
	}
	uid_int, err := strconv.Atoi(uid.(string))
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data.Chapter.UserID = int32(uid_int)

	new_chapter, err := ch.Chapter.CreateChapter(c, &data.Chapter)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	new_list_vocabulary := ch.Chapter.CreateVocabularyOfService(c, data.List_Vocabulary, pgtype.Int4{
		Int32: new_chapter.ID,
		Valid: true,
	})
	if len(new_list_vocabulary) == 0 {
		slog.Error("Error not vocabulary is created")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Error not vocabulary is created",
		})
		return
	}

	c.JSON(http.StatusOK, ChapterResponse{
		Chapter:         *new_chapter,
		List_Vocabulary: new_list_vocabulary,
	})

}
