package drumcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (d *deleteDrumController) Delete(c *gin.Context) {
	paramID := c.Params.ByName("id")

	ID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	err = d.usecase.Delete(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
