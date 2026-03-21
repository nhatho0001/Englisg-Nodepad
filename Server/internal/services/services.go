package services

import (
	"app-notepad/internal/store"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (u *UserService) getBase64FromHashToken(token *jwt.Token) (string, error) {
	h := sha256.New()
	h.Write([]byte(token.Raw))
	hashToken := h.Sum(nil)
	hash_refresh_token, err := bcrypt.GenerateFromPassword(hashToken, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Hash token is faild! \n Error : %s", err)
	}
	base64RefreshToken := base64.StdEncoding.EncodeToString([]byte(hash_refresh_token))
	return base64RefreshToken, nil
}

func (u *UserService) CreateToken(ctx context.Context, token *jwt.Token, uid int32) (*store.RefreshToken, error) {
	hash_refresh_token, err := u.getBase64FromHashToken(token)
	if err != nil {
		return nil, fmt.Errorf("Hash token is faild! \n Error : %s", err)
	}
	expire_date, err := token.Claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("Expire date is faild! Error : %v\n", err)
	}
	new_token, err := u.query.CreateToken(ctx, store.CreateTokenParams{
		UserID:      uid,
		HashedToken: hash_refresh_token,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ExpiresAt: pgtype.Timestamp{
			Time:  expire_date.Time,
			Valid: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("Create Refresh Token failt : %v\n", err)
	}

	return &new_token, nil

}

func (u *UserService) ByTokenAndUid(ctx context.Context, uid int32, token *jwt.Token) (*store.RefreshToken, error) {
	hash := sha256.Sum256([]byte(token.Raw))

	// var tokens []RefreshTokens
	tokens, err := u.query.GetTokensByUid(ctx, uid)

	if err != nil {
		return nil, err
	}

	for _, t := range tokens {
		stored, _ := base64.StdEncoding.DecodeString(t.HashedToken)

		if bcrypt.CompareHashAndPassword(stored, hash[:]) == nil {
			return &t, nil
		}
	}

	return nil, fmt.Errorf("Recode not found")
}

func (u *UserService) DeleteUserToken(ctx context.Context, uid int32) (bool, error) {
	if err := u.query.DeleteUserToken(ctx, uid); err != nil {
		return false, err
	}
	return true, nil
}
