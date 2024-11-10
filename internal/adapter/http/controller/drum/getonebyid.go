package drumcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (d *getOneByIDDrumController) GetOneByID(c *gin.Context) {
	paramID := c.Params.ByName("id")

	ID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	drum, err := d.usecase.GetOneByID(ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, drum)
}
