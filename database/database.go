package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/realjoni17/Cdlsiet/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//var host = ""
//var user = ""
//var pw = ""
//var port int
//var dbname = ""
//var ssl = ""

func StartDatabase() {
	host := os.Getenv("Host")
	user := os.Getenv("User")
	pw := os.Getenv("Password")
	port := os.Getenv("Port")
	dbName := os.Getenv("Db_Name")
	ssl := os.Getenv("SslMode")
	str := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, pw, dbName, port, ssl)
	//var cred = "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host,user,pw,dbName,port"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
