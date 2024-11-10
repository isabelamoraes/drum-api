package router

import (
	direview "drum-api/internal/di/http/review"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func useReviewRoute(conn gorm.DB, r *gin.Engine) {
	cc := direview.DICreateReviewController(&conn)

	r.POST("/v1/drum/:id/review", cc.Create)
}
