package db

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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

	err = createSchema(db)

	if err != nil {
		log.Fatal("Error creating schema")
	}

	Db = db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
