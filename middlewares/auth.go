package middlewares

import (
	"net/http"

	"github.com/AMCodeBytes/go-rest-film-review/utils"
	"github.com/gin-gonic/gin"
)

func Auth(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// If something goes wrong Abort will prevent anything after from running
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorised."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorised."})
		return
	}

	context.Set("userId", userId)

	// Call the next handler
	context.Next()
}
