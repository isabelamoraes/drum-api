package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func useRoute(r *gin.Engine) {
	r.GET("/v1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "API running"})
	})
}
