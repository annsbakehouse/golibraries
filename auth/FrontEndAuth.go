package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CollectionPageHitRules(c *gin.Context, endpoint string) (int, int) {
	frontend_hit := os.Getenv("frontend_hit")
	var arr []map[string]interface{}
	_ = json.Unmarshal([]byte(frontend_hit), &arr)
	path := endpoint
	for _, v := range arr {
		if v["path"] == path {
			s, _ := strconv.Atoi(fmt.Sprintf("%v", v["max_hit"]))
			ss, _ := strconv.Atoi(fmt.Sprintf("%v", v["time_range_second"]))
			return s, ss
		}
	}
	return 0, 0
}

type FrontEndTokenForm struct {
	Token string `json:"token" binding:"required"`
}

func ExtractFrontEndToken(c *gin.Context) (map[string]interface{}, error) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	fmt.Println(string(x))
	return nil, nil
}
func FrontEndTokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form FrontEndTokenForm
		if err := c.BindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
				"error":   true,
			})
			c.Abort()
			return
		}
		c.Next()
	}
	//var form FrontEndTokenForm
	// var form FrontEndTokenForm
	// if err := c.BindJSON(&form); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"message": err.Error(),
	// 		"error":   true,
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// return

	// if err := c.BindJSON(&form); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"message": err.Error(),
	// 		"error":   true,
	// 	})
	// 	c.Abort()
	// 	return
	// }

	// key := []byte(os.Getenv("chiper_token"))
	// text, _ := library.DecryptAES256(key, form.Token)
	// text = library.Base64Decode(text)
	// textSplit := strings.Split(text, "|")
	// if len(textSplit) != 5 {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"status":  http.StatusUnauthorized,
	// 		"message": "You have wrong authorization",
	// 		"error":   true,
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// clientip := c.ClientIP()
	// if textSplit[1] != clientip {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"status":  http.StatusUnauthorized,
	// 		"message": "You have wrong ip address",
	// 		"error":   true,
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// db, dbReader, _ := models.DbConnect()
	// defer db.Close()
	// max_hit, time_range_second := CollectionPageHitRules(c, textSplit[3])

	// if max_hit > 0 {
	// 	var model []models.HiltLogModel
	// 	tnow := time.Now()
	// 	then := time.Date(tnow.Year(), tnow.Month(), tnow.Day(), tnow.Hour(), tnow.Minute(), tnow.Second()-time_range_second, tnow.Nanosecond(), tnow.Location())
	// 	res := dbReader.Model(&model).Where("cust_time>=?", then.Format("2006-01-02 15:04:05"))
	// 	res.Where("cust_time<=?", tnow.Format("2006-01-02 15:04:05"))
	// 	res.Where("ip=?", textSplit[1])
	// 	res.Where("device_id=?", textSplit[0])
	// 	res.Find(&model)
	// 	if int(res.RowsAffected) >= max_hit {
	// 		c.JSON(http.StatusUnauthorized, gin.H{
	// 			"status":  http.StatusUnauthorized,
	// 			"message": "You have limit to access this page",
	// 			"error":   true,
	// 		})
	// 		c.Abort()
	// 	}
	// }
	// tx := dbReader.Begin()

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// 	db.Close()
	// }()

	// if err := tx.Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"message": err.Error,
	// 		"error":   true,
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// custTime, _ := time.Parse("2006-01-02 15:04:05", textSplit[2])
	// var model = models.HiltLogModel{
	// 	IP:         textSplit[1],
	// 	DeviceID:   textSplit[0],
	// 	CustTime:   custTime,
	// 	EndPoint:   textSplit[3],
	// 	DeviceInfo: textSplit[4],
	// }
	// result := tx.Create(&model)
	// if result.Error != nil {
	// 	tx.Rollback()
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"message": fmt.Sprintf("%v", result.Error) + "aaaa",
	// 		"error":   true,
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// tx.Commit()
}
