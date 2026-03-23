package controller

import (
	"app-notepad/internal/services"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserInput) ValidateInputData() bool {
	if u.Email == "" || u.Password == "" {
		return false
	}
	return true
}

type UserHander struct {
	Service *services.UserService
}

func NewUserHander(s *services.UserService) *UserHander {
	return &UserHander{
		Service: s,
	}
}

func (u *UserHander) UserLogin(c *gin.Context) {
	var user_login UserInput
	if err := c.ShouldBindBodyWithJSON(&user_login); err != nil {
		slog.Error(fmt.Sprintf("Parse Info User Is Faild %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	if !user_login.ValidateInputData() {
		slog.Error("Missing Email or Password")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing Email or Password"})
		return
	}

	current_user, err := u.Service.GetUser(c, user_login.Email)
	if err != nil {
		slog.Error(fmt.Sprintf("Error get User %v", err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found or invalid credentials"})
		return
	}

	if !u.Service.CheckPassword(c, user_login.Password, current_user.HashedPassword.String) {
		slog.Error("Password wrong")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Password wrong"})
		return
	}

	tokenPair, err := u.Service.GenerateJWT(strconv.FormatUint(uint64(current_user.ID), 10))

	if err != nil {
		slog.Error("Create Token Failed!")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tokens"})
		return
	}

	_, err = u.Service.DeleteUserToken(c, current_user.ID)
	if err != nil {
		slog.Error("Failed when cleaning Refresh Token!")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear old tokens"})
		return
	}

	_, err = u.Service.CreateToken(c, tokenPair.RefreshToken, current_user.ID)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to store new token"})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"AccesToken":   tokenPair.AcessToken.Raw,
		"RefreshToken": tokenPair.RefreshToken.Raw,
	})

}

func (u *UserHander) UserSignUp(c *gin.Context) {
	var user_signup UserInput
	if err := c.ShouldBindBodyWithJSON(&user_signup); err != nil {
		slog.Error("Parse data post is faild!")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Data post server is not parse!",
		})
		return
	}
	if !user_signup.ValidateInputData() {
		slog.Error("Email and Password is request")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Email and Password is request",
		})
		return
	}
	if _, err := u.Service.GetUser(c, user_signup.Email); err == nil {
		slog.Error("Email is exist")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Email is exist",
		})
		return
	}

	new_user, err := u.Service.CreateUserAccount(c, user_signup.Email, user_signup.Password)

	if err != nil {
		slog.Error("Recodre User Is not save")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Recodre User Is not save",
		})
		return
	}
	fmt.Print(new_user)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"email":    new_user.Email,
		"CreateAT": new_user.CreatedAt.Time,
	})

}
