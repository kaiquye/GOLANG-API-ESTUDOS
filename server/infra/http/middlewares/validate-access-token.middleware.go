package middlewares

import (
	"fmt"
	"fullVagas/infra/config"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthMiddleware struct {
	headersAccepted string
	tokenAccepted   string
	JWT             config.AccessTokenConfig
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		JWT:             *config.NewAccessTokenConfig(),
		headersAccepted: "authorization",
		tokenAccepted:   "Bearer",
	}
}

func (middleware *AuthMiddleware) ValidateAccessToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		accessToken := context.GetHeader(middleware.headersAccepted)
		if len(accessToken) <= 0 {
			context.JSON(401, gin.H{
				"message": "not authorized",
			})
			context.Abort()
			return
		}

		tokenSlice := strings.Split(accessToken, " ")
		if len(tokenSlice) != 2 || tokenSlice[0] != middleware.tokenAccepted {
			context.JSON(401, gin.H{
				"message": "not authorized",
			})
			context.Abort()
			return
		}

		valid, err := middleware.JWT.Validate(tokenSlice[1])
		fmt.Println("tested")
		if err != nil || *valid == false {
			context.JSON(401, gin.H{
				"message": "not authorized",
			})
			context.Abort()
			return
		}

		context.Next()
		return
	}
}
