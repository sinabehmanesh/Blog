package database_contact

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// ///////////////
// Types
// ///////////////
type Contact struct {
	gorm.Model
	email   string
	subject string
	message string
	DATE    string
}

type Contacts []Contact

func Insert_message(contactemail string, contactsubject string, contactmessage string) {

	dt := time.Now()

	// printing the time in string format
	nowdate := dt.Format(time.RFC3339)

	var record Contact
	record.email = contactemail
	record.subject = contactsubject
	record.message = contactmessage
	record.DATE = nowdate

	fmt.Println(record.email, record.message, record.subject)
	fmt.Println("view check!")
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

	db, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()

	query := `INSERT INTO contacts (email, subject, message, date) VALUES (` + record.email + `,` + record.subject + "," + record.message + "," + nowdate + ");"

	fmt.Println(query)
}
