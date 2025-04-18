package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prezessikora/events/utils"
)

func Authenticate(ctx *gin.Context) {
	//  authenticate the route
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		fmt.Println(err)
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()

}
