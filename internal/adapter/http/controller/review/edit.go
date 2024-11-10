package reviewcontroller

import (
	"drum-api/internal/core/dto"
	"drum-api/internal/infra/gorm/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *reviewController) Create(c *gin.Context) {
	paramID := c.Params.ByName("id")

	ID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	var bodyReq dto.CreateReviewRequest

	if err := c.ShouldBindJSON(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	review := model.Review{
		Rating:  bodyReq.Rating,
		Comment: bodyReq.Comment,
		DrumID:  ID,
	}

	err = r.usecase.Create(&review)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
