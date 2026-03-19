package services

import (
	"app-notepad/internal/store"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	query *store.Queries
}

func NewUserService(q *store.Queries) *UserService {
	return &UserService{query: q}
}

func (u *UserService) GetUser(ctx context.Context, email string) (*store.User, error) {
	current_user, err := u.query.GetAuthor(ctx, email)
	if err != nil {
		return nil, err
	}
	return &current_user, nil
}

func (u *UserService) CheckPassword(ctx context.Context, password string, hash_password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password)); err != nil {
		return false
	}
	return true
}
