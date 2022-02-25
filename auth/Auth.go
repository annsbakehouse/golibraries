package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"annsbakehouse.com/erp1/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ExtractToken(c *gin.Context) string {
	//getting token from bearer header
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

//getting token to claims
func ExtractTokenID(c *gin.Context) (map[string]interface{}, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("secretKey")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims,nil
	}
	return nil, nil
}

//token validation
func TokenValid(c *gin.Context) {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		_,dbReader,err := models.DbConnect()
		if err!=nil {
			return nil, fmt.Errorf("Signing Token Expired")
		}
		tt := time.Now()
		tFormat := tt.Format("2006-01-02 15:04:05")
		var activeUserModel models.ActiveUserModel
		recActiveUser := dbReader.Select("id").First(&activeUserModel,"token=? AND expired_on>?",tokenString,tFormat)
		if recActiveUser.Error==nil && recActiveUser.RowsAffected==1 {
			return []byte(os.Getenv("secretKey")), nil
		}
		return nil, fmt.Errorf("Signing Token Expired")
	})
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Your have wrong Authorization",
			"error":true,
		})
		c.Abort()
	}
}