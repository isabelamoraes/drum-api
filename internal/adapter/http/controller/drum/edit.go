package drumcontroller

import (
	"drum-api/internal/core/dto"
	"drum-api/internal/infra/gorm/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (d *editDrumController) Edit(c *gin.Context) {
	paramID := c.Params.ByName("id")

	ID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

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

	err = d.usecase.Edit(ID, &drum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
