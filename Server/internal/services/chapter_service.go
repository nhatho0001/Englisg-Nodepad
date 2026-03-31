package services

import (
	"app-notepad/configs"
	"app-notepad/internal/store"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ChapterService struct {
	cfg   *configs.Configs
	query *store.Queries
}

type ChapterInfo struct {
	Chapter         *store.Chapter
	List_Vocabulary []store.Vocabulary
}

type ChapterInfoParams struct {
	Chapter         *store.UpdateChaptersParams
	List_Vocabulary []store.CreateVocabularyParams
}

func (cp *ChapterInfoParams) ValidateDataInput() bool {
	if !cp.Chapter.Title.Valid {
		return false
	}
	return true
}

func NewChapterService(q *store.Queries, cfg *configs.Configs) *ChapterService {
	return &ChapterService{query: q, cfg: cfg}
}

func (chapter *ChapterService) GetChapterById(ctx context.Context, id int32) (*store.Chapter, error) {
	data, err := chapter.query.GetChaptersById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (chapter *ChapterService) GetVocabularyOfChapter(ctx context.Context, chapter_id pgtype.Int4) ([]store.Vocabulary, error) {
	data, err := chapter.query.GetVocabularyOfChapter(ctx, chapter_id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (chapter *ChapterService) GetListChapter(ctx context.Context, uid int32) ([]store.Chapter, error) {
	list_chapter, err := chapter.query.GetChaptersByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return list_chapter, nil
}

func (chapter *ChapterService) CreateChapter(ctx context.Context, arg *store.CreateChapterParams) (*store.Chapter, error) {
	if !arg.Status.Valid {
		arg.Status.String = "new"
	}
	fmt.Println(arg)
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

func (chapter *ChapterService) UpdateChapterService(ctx context.Context, arg store.UpdateChaptersParams) (*store.Chapter, error) {
	chapter_update, err := chapter.query.UpdateChapters(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &chapter_update, nil
}

func (chapter *ChapterService) UpdateVocabularyService(ctx context.Context, arg []store.UpdateVocabularyParams) ([]store.Vocabulary, error) {
	vocabulary_list := make([]store.Vocabulary, len(arg))
	for index, vocabulary := range arg {
		update_vocabulary, err := chapter.query.UpdateVocabulary(ctx, vocabulary)
		if err == nil {
			vocabulary_list[index] = update_vocabulary
		}
	}
	return vocabulary_list, nil
}

func (chapter *ChapterService) UpdateChapterAndVocabulary(ctx context.Context, db *pgx.Conn, arg *ChapterInfoParams) (*ChapterInfo, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	update_chapter, err := chapter.UpdateChapterService(ctx, *arg.Chapter)
	if err != nil {
		return nil, err
	}
	if err := chapter.query.DeleteVocabularyOfChapter(ctx, pgtype.Int4{
		Int32: update_chapter.ID,
		Valid: true,
	}); err != nil {
		return nil, err
	}
	list_update_vocabulary := chapter.CreateVocabularyOfService(ctx, arg.List_Vocabulary, pgtype.Int4{
		Int32: update_chapter.ID,
		Valid: true,
	})
	tx.Commit(ctx)
	return &ChapterInfo{
		Chapter:         update_chapter,
		List_Vocabulary: list_update_vocabulary,
	}, nil
}

func (chapter *ChapterService) DeleteChapter(ctx context.Context, db *pgx.Conn, chapter_id int32) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err := chapter.query.DeleteVocabularyOfChapter(ctx, pgtype.Int4{
		Int32: chapter_id,
		Valid: true,
	}); err != nil {
		return err
	}
	if err := chapter.query.DeleteChapter(ctx, chapter_id); err != nil {
		return err
	}
	tx.Commit(ctx)

	return nil
}
