package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	log.Default().Print("Connecting to database...")
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABES_CONNECTION_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Default().Print("Connected to database")
	db.Migrator().DropTable("links")
	db.Migrator().CreateTable(&Link{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&Link{})

	DB = db
}
