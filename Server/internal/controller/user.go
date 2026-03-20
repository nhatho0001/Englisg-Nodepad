package controller

import (
	"app-notepad/internal/services"
	"fmt"
	"log/slog"
	"net/http"

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
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if !user_login.ValidateInputData() {
		slog.Error("Missing Email or Password")
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Missing Info Email or Password"))
		return
	}

	current_user, err := u.Service.GetUser(c, user_login.Email)
	if err != nil {
		slog.Error(fmt.Sprintf("Error get User %v", err))
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if !u.Service.CheckPassword(c, user_login.Password, current_user.HashedPassword.String) {
		slog.Error("Password wrong")
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Password wrong"))
		return
	}
	fmt.Print(user_login)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"email":    user_login.Email,
		"password": user_login.Password,
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
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"email":    new_user.Email,
		"CreateAT": new_user.CreatedAt.Time,
	})

}
