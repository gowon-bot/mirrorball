package db

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

// Db holds a database reference
var Db *pg.DB

// InitDB initializes the database
func InitDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	parsedOptions, err := pg.ParseURL(databaseURL)

	if err != nil {
		log.Fatal("Error parsing database url")
	}

	db := pg.Connect(parsedOptions)

	db.AddQueryHook(Logger{})

	Db = db
}
