package library

import (
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"strconv"
)

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
)

func SaveImageToDisk(fileNameBase, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidImage
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}
	imgCfg, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	if imgCfg.Width != 750 || imgCfg.Height != 685 {
		// return "", ErrSize
	}

	fileName := fileNameBase + "." + fm
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)

	return fileName, err
}

func GenUID(sequence int, lengthNumber int, charMinLenght int) string {
	stringSequence := fmt.Sprintf("%v", sequence)
	if len(stringSequence) > lengthNumber {
		fmt.Println(stringSequence)
		prefix := stringSequence[0 : len(stringSequence)-lengthNumber]
		suffix := stringSequence[len(stringSequence)-lengthNumber : len(stringSequence)]
		numIn, _ := strconv.Atoi(prefix)
		output := GenUIDToAlpha(numIn, "")
		if len(output) < charMinLenght {
			iOutput := charMinLenght - len(output)
			n := 0
			for n < iOutput {
				output = "A" + output
				n++
			}
		}
		return output + suffix
	} else {

		if len(stringSequence) < lengthNumber {
			i := 0
			for i < lengthNumber-len(stringSequence) {
				stringSequence = "0" + stringSequence
				i++
			}
		}
		n := 0
		for n < charMinLenght {
			stringSequence = "A" + stringSequence
			n++
		}
		return stringSequence
	}
	return ""
}
func GenUIDToAlpha(num int, addString string) string {
	if num <= 26 {
		return string('A'-1+num) + addString
	} else {
		total := num - 26
		return GenUIDToAlpha(total, string('A'-1+26)+addString)
	}
}

//json encode
func JsonEncode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}

//json decode
func JsonDecode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	json.Unmarshal([]byte(data), &dat)
	return dat, nil
}

//parse json requst Body
func JsonReqBody(c *gin.Context) (map[string]interface{}, error) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var dat map[string]interface{}
	json.Unmarshal([]byte(value), &dat)
	return dat, nil
}

func IsDateValue(stringDate string) bool {
	_, err := time.Parse("2006-01-02", stringDate)
	return err == nil
}
