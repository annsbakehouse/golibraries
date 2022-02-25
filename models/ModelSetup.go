package models
import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	//arango "github.com/joselitofilho/gorm-arango/pkg"
	// "github.com/arangodb/go-driver/http"
	// driver "github.com/arangodb/go-driver"
	"os"
	"gorm.io/gorm/logger"
	//arango "github.com/arangodb/go-driver"
	// "github.com/arangodb/go-driver/http"
	"fmt"
	// "github.com/fatih/structs"
	// "strings"
	// "os"
	"encoding/json"
	// "reflect"
	"database/sql"
	"time"
)

func DbConnect() (*gorm.DB,*gorm.DB,error) {
	//mysql connection
	configDbMaster := os.Getenv("masterDsn");
	dbMaster, err := gorm.Open(postgres.Open(configDbMaster),&gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
		QueryFields: true,
	})
	if err != nil {
		return nil,nil,fmt.Errorf("Master Database Connection Error")
	}
	//end mysql connection

	//arango connection
	// conn, err := http.NewConnection(http.ConnectionConfig{
	// 	Endpoints: []string{os.Getenv("golangURI")},
	// })
	// if err != nil {
	// 	log.Panic(err)
	// }

	// // Client object
	// client, err := driver.NewClient(driver.ClientConfig{
	// 	Connection: conn,
	// 	Authentication: driver.BasicAuthentication(os.Getenv("golangUser"), os.Getenv("golangPassword")),
	// })
	// if err != nil {
	// 	log.Panic(err)
	// }

	// // Open "examples_books" database
	// _, err := client.Database(nil, os.Getenv("golangDatabase"))
	// if err != nil {
	// 	log.Panic(err)
	// }
	//end arango connections
	return dbMaster,dbMaster,nil
}

//json encode
// func JsonEncode(data interface{}) (string, error) {
// 	jsons, err := json.Marshal(data)
// 	return string(jsons), err
// }

// //json decode
// func JsonDecode(data string) (map[string]interface{}, error) {
// 	var dat map[string]interface{}
// 	json.Unmarshal([]byte(data),&dat)
// 	return dat, nil
// }


//parse null string on model
type NullString struct {
	sql.NullString
}
//parse null string on model
func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return []byte(`null`),nil
}

func NullStringInput(s string) (NullString) {
	
	return NullString{sql.NullString{s,true}}
}

//parse null int on model
type NullInt64 struct {
	sql.NullInt64
}

//parse null int on model
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return []byte(`null`),nil
}

//parse null time on model
type NullDateTime struct {
	sql.NullString
}
//parse null time on model
func (nt NullDateTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		t, _ := time.Parse(time.RFC3339, nt.String)
		return json.Marshal(t.Format("2006-01-02 15:04:05"))
	}
	return []byte(`null`),nil
}

//get Columns Name from Model
func GetColumnName(db *gorm.DB, model []interface{}) (*gorm.DB){
	for _,v := range model {
		fmt.Println(v)
	}
	return db
}