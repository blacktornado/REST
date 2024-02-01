package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Printf("Error Loading .env File: %s\n", err.Error())
		return
	}

	connInfo := connection{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DB"),
	}

	db, err := sql.Open("mysql", connToString(connInfo))
	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB pinged successfully\n")

	}
	DB = db
	//defer db.Close()

}

func connToString(info connection) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", info.User, info.Password, info.Host, info.Port, info.DBName)
}
