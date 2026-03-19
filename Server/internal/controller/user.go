package controller

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserLogin(c *gin.Context) {
	var user_login UserRequest
	if err := c.ShouldBindBodyWithJSON(&user_login); err != nil {
		slog.Error(fmt.Sprintf("Parse Info User Is Faild %v", err))
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Print(user_login)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"email":    user_login.Email,
		"password": user_login.Password,
	})

}
