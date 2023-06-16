package library

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"text/template"

	"github.com/arangodb/go-driver"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func SyncTelegramM2Bot(order_data_id string, db driver.Database) {

	b, err := os.ReadFile(os.Getenv("upload_path_view") + "query/arango/botTelegram.sql") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// bot, err := tgbotapi.NewBotAPI("6032714656:AAEJigZalKyz1KW_q9BLHYftWFdbCvgcnR4")
	bot, err := tgbotapi.NewBotAPI(os.Getenv("telegram_bot_dev"))
	if err != nil {
		log.Fatal(err.Error())
	}
	bot.Debug = true
	tx := context.Background()
	var filepath = path.Join(os.Getenv("upload_path_view"), "email_template", "telegram_paid.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	cursor, err := db.Query(tx, string(b), map[string]interface{}{"id": order_data_id})
	if err != nil {
		log.Panic(err.Error())
		return
	}
	var data map[string]interface{}
	for {
		var doc map[string]interface{}
		_, err := cursor.ReadDocument(nil, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatalf("Doc returned: %v", err.Error())
		} else {

		}
		data = doc
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		log.Panic(err.Error())
		return
	}
	result := tpl.String()
	i, _ := strconv.ParseInt(os.Getenv("telegram_channel_order_paid_message"), 10, 64)
	msg := tgbotapi.NewMessage(i, result)
	bot.Send(msg)
}
func SyncTelegramM2BotChecklist(order_data_id string, db driver.Database, dbdef *gorm.DB) {

	b, err := os.ReadFile(os.Getenv("upload_path_view") + "query/arango/botTelegramChecklist.sql") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("telegram_bot_dev"))
	if err != nil {
		log.Fatal(err.Error())
	}
	bot.Debug = true
	tx := context.Background()
	var filepath = path.Join(os.Getenv("upload_path_view"), "email_template", "botTelegramChecklist.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	cursor, err := db.Query(tx, string(b), map[string]interface{}{"id": order_data_id})
	if err != nil {
		log.Panic(err.Error())
		return
	}
	var data map[string]interface{}
	for {
		var doc map[string]interface{}
		_, err := cursor.ReadDocument(nil, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatalf("Doc returned: %v", err.Error())
		} else {

		}
		data = doc
		data_warehouse := make(map[string]interface{})
		dbdef.Raw("Select * FROM warehouse where id=?", data["wh_id"]).Scan(&data_warehouse)
		fmt.Println(data_warehouse)
		data["wh_name"] = fmt.Sprintf("%v (%v)", data_warehouse["name"], data_warehouse["letter_code"])
		created_by := fmt.Sprintf("%v", data["created_by"])
		if len(created_by) == 0 {
			created_by = "Website (" + fmt.Sprintf("%v", data["customer_name"]) + ")"
		} else {
			data_users := make(map[string]interface{})
			dbdef.Raw("Select * FROM users WHERE id=?", created_by).Scan(&data_users)
			created_by = fmt.Sprintf("%v (%v)", data_users["UID"], data_users["fullname"])
		}
		data["created_by"] = created_by
		delivery_fee, _ := data["delivery_fee"].(float64)
		data["delivery_fee"] = NumberFormat(delivery_fee, 0, "", ",")
		total_price_discount, _ := data["total_price_discount"].(float64)
		data["total_price_discount"] = NumberFormat(total_price_discount, 0, "", ",")
		grand_total, _ := data["grand_total"].(float64)
		data["grand_total"] = NumberFormat(grand_total, 0, "", ",")
	}
	// fmt.Println(data)
	// return
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		log.Panic(err.Error())
		return
	}
	result := tpl.String()
	i, _ := strconv.ParseInt(os.Getenv("telegram_channel_Checklist_message"), 10, 64)
	msg := tgbotapi.NewMessage(i, result)
	bot.Send(msg)
}
