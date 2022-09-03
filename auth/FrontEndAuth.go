package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CollectionPageHitRules(c *gin.Context) int {
	frontend_hit := os.Getenv("frontend_hit")
	var arr []map[string]interface{}
	_ = json.Unmarshal([]byte(frontend_hit), &arr)
	path := c.FullPath()
	for _, v := range arr {
		if v["path"] == path {
			return v["max_hit"].(int)
		}
	}
	return 0
}

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

	// max_hit := CollectionPageHitRules(c)
	// fmt.Println(max_hit)
	return
}
