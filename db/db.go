package db

import (
	"dungeons-and-dragons/db/seeders"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		os.Getenv("DB_PASSWORD"),
		host,
		port,
		os.Getenv("DB_NAME"),
	)

	log.Printf("Connecting to %v on port %v with username %v", host, port, user)
	db, err = gorm.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err.Error())
	}

	seeder := seeders.NewSeeder(db)
	seeder.SetClasses()
	seeder.SetRaces()
	seeder.SetCharacters()

	return db
}

func DB() *gorm.DB {
	return db
}
