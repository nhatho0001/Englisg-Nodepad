package services

import (
	"app-notepad/configs"
	"app-notepad/internal/store"
	"context"
)

type ChapterService struct {
	cfg   *configs.Configs
	query *store.Queries
}

func NewChapterService(q *store.Queries, cfg *configs.Configs) *ChapterService {
	return &ChapterService{query: q, cfg: cfg}
}

func (chepter *ChapterService) GetListChapter(ctx context.Context, uid int32) ([]store.Chapter, error) {
	list_chapter, err := chepter.query.GetChaptersByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return list_chapter, nil
}
