package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"url-shortener/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB {
	if os.Getenv("RENDER") == "" {
		// means it is being run in local
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading env variables")
		}
	}

	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DB := os.Getenv("DB_NAME")
	PORT := os.Getenv("DB_PORT")

	// dsn := USER + ":" + PASSWORD + "@(" + HOST + PORT + ")/" + DB
	dsn := fmt.Sprintf("%s:%s@(%s%s)/%s", USER, PASSWORD, HOST, PORT, DB)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting with database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error while trying to ping DB:", err)
	}

	// err = models.CreateURLTable()
	err = models.CreateURLTable(db)
	if err != nil {
		log.Fatal("Error while creating table:", err)
	}

	return db
}
