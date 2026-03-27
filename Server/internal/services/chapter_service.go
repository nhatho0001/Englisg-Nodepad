package services

import (
	"app-notepad/configs"
	"app-notepad/internal/store"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type ChapterService struct {
	cfg   *configs.Configs
	query *store.Queries
}

func NewChapterService(q *store.Queries, cfg *configs.Configs) *ChapterService {
	return &ChapterService{query: q, cfg: cfg}
}

func (chapter *ChapterService) GetListChapter(ctx context.Context, uid int32) ([]store.Chapter, error) {
	list_chapter, err := chapter.query.GetChaptersByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return list_chapter, nil
}

func (chapter *ChapterService) CreateChapter(ctx context.Context, arg *store.CreateChapterParams) (*store.Chapter, error) {
	if arg.Status.String == "" {
		arg.Status.String = "new"
	}
	new_chapter, err := chapter.query.CreateChapter(ctx, *arg)
	if err != nil {
		return nil, err
	}
	return &new_chapter, nil
}

func (chapter *ChapterService) CreateVocabularyOfService(ctx context.Context, arg []store.CreateVocabularyParams, chapeter_id pgtype.Int4) []store.Vocabulary {
	var result []store.Vocabulary
	for _, v := range arg {
		v.ChapterID = chapeter_id
		new_vocabulary, err := chapter.query.CreateVocabulary(ctx, v)
		if err == nil {
			result = append(result, new_vocabulary)
		}
	}
	return result
}
