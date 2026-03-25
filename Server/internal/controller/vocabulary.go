package controller

import (
	"app-notepad/internal/services"
	"app-notepad/internal/store"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type VocabularyHandler struct {
	vocabulary_service *services.VocabularyService
}

type SearchFilters struct {
	ID     int32  `form:"id"`
	Query  string `form:"q"`
	Limit  int32  `form:"limit"`
	Offset int32  `form:"offset"`
}

func NewVocabularyHandler(vs *services.VocabularyService) *VocabularyHandler {
	return &VocabularyHandler{
		vocabulary_service: vs,
	}
}

func (vs *VocabularyHandler) CreateVocabulary(c *gin.Context) {
	var input_vocabulary store.CreateVocabularyParams
	if err := c.ShouldBindBodyWithJSON(&input_vocabulary); err != nil {
		slog.Error("Parse data post is faild!")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Data post server is not parse!",
		})
		return
	}

	new_vocabulary, err := vs.vocabulary_service.CreateVocabulary(c, input_vocabulary)
	if err != nil {
		slog.Error("Save vocabulary is failed!")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Save vocabulary is failed!",
		})
		return
	}

	c.JSON(http.StatusOK, new_vocabulary)
}

func (vs *VocabularyHandler) GetVocabularyOfChapter(c *gin.Context) {
	var query_data SearchFilters
	if err := c.BindQuery(&query_data); err != nil {
		slog.Error("Parse data post is faild!")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Data post server is not parse!",
		})
		return
	}

	fmt.Println(query_data, query_data.ID, query_data.Limit, query_data.Offset)

	if query_data.Limit == 0 {
		query_data.Limit = 20
	}

	list_vocabulary, err := vs.vocabulary_service.GetVocabularyOfChapter(c, &store.GetCharacterVocabularyParams{
		ChapterID: pgtype.Int4{
			Int32: query_data.ID,
			Valid: true,
		},
		Limit:  query_data.Limit,
		Offset: query_data.Offset,
	})

	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Get vocabulary is failed!",
		})
		return
	}

	c.JSON(http.StatusOK, list_vocabulary)
}

func (vs *VocabularyHandler) GetVocabularyOfUser(c *gin.Context) {
	var query_data SearchFilters
	if err := c.BindQuery(&query_data); err != nil {
		slog.Error("Parse data post is faild!")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Data post server is not parse!",
		})
		return
	}

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

	if query_data.Limit == 0 {
		query_data.Limit = 20
	}

	list_vocabulary, err := vs.vocabulary_service.GetVocabularyOfUser(c, &store.GetVocabularyOfUserParams{
		UserID: int32(uidInt),
		Limit:  query_data.Limit,
		Offset: query_data.Offset,
	})

	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Get vocabulary is failed!",
		})
		return
	}

	c.JSON(http.StatusOK, list_vocabulary)

}
