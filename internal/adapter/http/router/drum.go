package router

import (
	didrum "drum-api/internal/di/http/drum"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func useDrumRoute(conn gorm.DB, r *gin.Engine) {
	cc := didrum.DICreateDrumController(&conn)
	lc := didrum.DIListDrumController(&conn)
	gc := didrum.DIGetOneByIDDrumController(&conn)
	dc := didrum.DIDeleteDrumController(&conn)
	ec := didrum.DIEditDrumController(&conn)

	r.POST("/v1/drum", cc.Create)
	r.GET("/v1/drum", lc.List)
	r.GET("/v1/drum/:id", gc.GetOneByID)
	r.DELETE("/v1/drum/:id", dc.Delete)
	r.PUT("/v1/drum/:id", ec.Edit)
}
