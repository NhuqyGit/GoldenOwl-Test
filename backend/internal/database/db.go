package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func ConnPostGresDB() *gorm.DB {
	once.Do(func() {
		var err error
		dbHost := os.Getenv("DB_HOST")
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbPort := os.Getenv("DB_PORT")

		connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPass, dbName, dbPort)
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Fatal("Database connection failed:", err)
		}

		log.Println("Connected to Database!")

		Migrate(db)
	})

	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		return ConnPostGresDB()
	}
	return db
}
