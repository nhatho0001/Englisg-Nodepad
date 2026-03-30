package services

import (
	"app-notepad/configs"
	"app-notepad/internal/store"
	"context"
)

type VocabularyService struct {
	cfg   *configs.Configs
	query *store.Queries
}

func NewVocabularyService(cfg *configs.Configs, query *store.Queries) *VocabularyService {
	return &VocabularyService{
		cfg:   cfg,
		query: query,
	}
}

func (v *VocabularyService) GetVocabularyOfChapter(ctx context.Context, arg *store.GetCharacterVocabularyParams) ([]store.Vocabulary, error) {
	list_character, err := v.query.GetCharacterVocabulary(ctx, *arg)
	if err != nil {
		return nil, err
	}
	return list_character, nil
}

func (v *VocabularyService) GetVocabularyOfUser(ctx context.Context, arg *store.GetVocabularyOfUserParams) ([]store.GetVocabularyOfUserRow, error) {
	list_character, err := v.query.GetVocabularyOfUser(ctx, *arg)
	if err != nil {
		return nil, err
	}
	return list_character, nil
}

func (v *VocabularyService) CreateVocabulary(ctx context.Context, arg store.CreateVocabularyParams) (*store.Vocabulary, error) {
	vocabulary, err := v.query.CreateVocabulary(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &vocabulary, nil
}
