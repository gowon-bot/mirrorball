package db

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Db holds a database reference
var Db *pg.DB

// InitDB initializes the database
func InitDB() {
	databaseURL := os.Getenv("DATABASE_URL")

	parsedOptions, err := pg.ParseURL(databaseURL)

	if err != nil {
		log.Fatal("Error parsing database url")
	}

	orm.RegisterTable((*RateYourMusicAlbumAlbum)(nil))
	orm.RegisterTable((*ArtistTag)(nil))

	db := pg.Connect(parsedOptions)

	db.AddQueryHook(Logger{})

	Db = db
}
