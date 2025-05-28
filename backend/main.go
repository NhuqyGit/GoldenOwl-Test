package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"goldenowl-test/internal/database"
	"goldenowl-test/internal/handlers"
	"goldenowl-test/internal/repositories"
	"goldenowl-test/internal/routers"
	"goldenowl-test/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main() {
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
	}

	// 1. Init & Conn DB
	dbProvider := database.NewDatabase()
	db := dbProvider.GetDB()
	pgxPool := dbProvider.GetPgxPool()

	// 2. New 3 layer
	scoreRepo := repositories.NewStudentScoreRepo(db, pgxPool)
	scoreService := services.NewStudentScoreService(scoreRepo)
	scoreHandler := handlers.NewStudentScoreHandler(scoreService)

	// 3. Seeder
	start := time.Now() // bắt đầu đếm thời gian

	if err := scoreService.SeedStudentScores(); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}else{
		log.Println("Seeding succeed")
	}

	duration := time.Since(start) // tính thời gian đã trôi qua
	fmt.Println("Time:", duration)

	r := gin.Default()
	routers.RegisterSubjectRoutes(r, scoreHandler)

	r.Run(":" + os.Getenv("PORT"))
}