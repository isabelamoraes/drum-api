package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(conn gorm.DB) *gin.Engine {
	r := gin.Default()

	useDrumRoute(conn, r)
	useReviewRoute(conn, r)
	useRoute(r)
	return r
}
