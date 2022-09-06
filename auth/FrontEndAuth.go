package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/annsbakehouse/golibraries/library"
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

func FrontEndTokenCheck(c *gin.Context) {
	// body := c.Request.Body
	// t, _ := ioutil.ReadAll(body)
	// stringJson, _ := library.JsonDecode(string(t))
	// if stringJson["token"] == nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"message": "Token Request Required",
	// 		"error":   true,
	// 	})
	// 	c.Abort()
	// }
	// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(t))

	// token := fmt.Sprintf("%v", stringJson["token"])
	// key := []byte(os.Getenv("chiper_token"))
	// text, _ := library.DecryptAES256(key, token)
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

	// 	res := dbReader.Model(&model).Where("created_at>=?", then.Format("2006-01-02 15:04:05"))
	// 	res.Where("ip=?", textSplit[1])
	// 	res.Where("device_id=?", textSplit[0])
	// 	res.Find(&model)
	// 	fmt.Println(res.RowsAffected, " ", max_hit)
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
	// 		"message": fmt.Sprintf("%v", result.Error),
	// 		"error":   true,
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// tx.Commit()
}

type GlobalToken struct {
	UID        string `json:"uid" binding:"required"`
	EndPoint   string `json:"endpoint"`
	DeviceInfo string `json:"info"`
}

func GetFrontendGlobalToken(c *gin.Context) {
	var form GlobalToken
	if err := c.BindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	StringToken := form.UID + "|" + c.ClientIP() + "|" + time.Now().Format("2006-01-02 15:04:05") + "|" + form.EndPoint + "|" + form.DeviceInfo
	StringToken = library.Base64Encode(StringToken)
	key := []byte(os.Getenv("chiper_token"))
	StringToken, _ = library.EncryptAES256(key, StringToken)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Proccess Done",
		"token":   StringToken,
		"error":   false,
	})
	return
}

type GlobalTokenCheck struct {
	Token string `json:"token" binding:"required"`
}

func GetFrontendCheckToken(c *gin.Context) {
	var form GlobalTokenCheck
	if err := c.BindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	token := form.Token
	key := []byte(os.Getenv("chiper_token"))
	text, _ := library.DecryptAES256(key, token)
	text = library.Base64Decode(text)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Proccess Done",
		"token":   text,
		"error":   false,
	})
	return
}
