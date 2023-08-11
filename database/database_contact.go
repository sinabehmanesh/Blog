package database_contact

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ///////////////
// Types
// ///////////////
type contact struct {
	email   string
	subject string
	message string
	date    string
}

func Insert_message(email string, subject string, message string) {

	dt := time.Now()

	// printing the time in string format
	nowdate := dt.Format(time.RFC3339)

	contactmessage := contact{
		email:   email,
		subject: subject,
		message: message,
		date:    nowdate,
	}

	///////////////////////
	// Database credentials
	///////////////////////
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can not load .env file!! Err: %s", err)
	}

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := "blog"

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	_ = db

	resault := db.Create(&contactmessage)

	fmt.Println(resault.Statement)

}
