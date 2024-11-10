package drumcontroller

import (
	"drum-api/internal/core/dto"
	"drum-api/internal/infra/gorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *createDrumController) Create(c *gin.Context) {
	var bodyReq dto.UpsertDrumRequest

	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	drum := model.Drum{
		Name:          bodyReq.Name,
		Mark:          bodyReq.Mark,
		Configuration: bodyReq.Configuration,
		IsEletronic:   *bodyReq.IsEletronic,
	}

	err := d.usecase.Create(&drum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
