package services

import (
	"app-notepad/internal/store"
	"context"

	"github.com/jackc/pgx/v5/pgtype"

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

func (u *UserService) CreateUserAccount(ctx context.Context, email string, password string) (*store.User, error) {
	hash_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	new_user := store.CreateUserParams{
		Email: email,
		HashedPassword: pgtype.Text{
			String: string(hash_password),
			Valid:  true,
		},
	}
	user, err := u.query.CreateUser(ctx, new_user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
