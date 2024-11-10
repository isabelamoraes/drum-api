package drumcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *listDrumController) List(c *gin.Context) {
	drums, err := d.usecase.List()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, drums)
}
