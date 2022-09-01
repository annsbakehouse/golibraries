package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FrontEndTokenForm struct {
	Token string `json:"token" binding:"required"`
}

func FrontEndTokenCheck(c *gin.Context) {
	var form FrontEndTokenForm
	if err := c.BindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	return
}
