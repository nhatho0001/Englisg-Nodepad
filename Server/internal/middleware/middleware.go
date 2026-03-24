package middleware

import (
	"app-notepad/configs"
	"app-notepad/internal/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomMiddleware struct {
	Service *services.UserService
	cfg     *configs.Configs
}

func NewMiddleware(service *services.UserService) *CustomMiddleware {
	return &CustomMiddleware{
		Service: service,
	}
}
func (m *CustomMiddleware) NewAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token := c.GetHeader("Authorization")
		if access_token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Not Accesstoken"})
			return
		}
		parseToken, err := m.Service.Parse(access_token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Acess Token is faild!"})
			return
		}
		uid, err := parseToken.Claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Parse Token can't get data neccessary"})
			return
		}
		expire_date, err := parseToken.Claims.GetExpirationTime()
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		currentTime := time.Now()
		if currentTime.After(expire_date.Time) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Access token is expired"})
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}

func (m *CustomMiddleware) AuthenOwnerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uidAny, exist := ctx.Get("uid")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You cannot access this page!"})
			return
		}

		uidStr, ok := uidAny.(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid user ID format in context"})
			return
		}

		uidInt, err := strconv.Atoi(uidStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: invalid user ID"})
			return
		}

		uidParamStr := ctx.Param("uid")
		uidParamInt, err := strconv.Atoi(uidParamStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Parameter is not suitable"})
			return
		}

		if uidInt != uidParamInt {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		ctx.Next()
	}
}
