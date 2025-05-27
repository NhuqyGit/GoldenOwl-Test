// internal/db/migrate.go
package database

import (
	"goldenowl-test/internal/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.StudentScore{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration succeed")
}

