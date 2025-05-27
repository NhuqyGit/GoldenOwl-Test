package main

import (
	"fmt"
	"log"
	"time"

	"goldenowl-test/internal/database"

	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 1. Kết nối DB
	db := database.ConnPostGresDB()

	// 2. Migration
	database.Migrate(db)

	// 3. Seeder
	start := time.Now() // bắt đầu đếm thời gian

	if err := database.SeedStudentScores(db); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}else{
		log.Println("Seeding succeed")
	}

	duration := time.Since(start) // tính thời gian đã trôi qua
	fmt.Println("Thời gian chạy:", duration)
}