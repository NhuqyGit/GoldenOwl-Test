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

	// Swagger docs import
	_ "goldenowl-test/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           GoldenOwl Student Score API
// @version         1.0
// @description     This is an API for student score management.
// @host            go-backend-rc59.onrender.com
// @BasePath        /
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
	start := time.Now()
	if err := scoreService.SeedStudentScores(); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}else{
		log.Println("Seeding succeeded")
	}
	fmt.Println("Time:", time.Since(start))

	// 4: Set up Gin server
	r := gin.Default()

	// 5: Register API routes
	routers.RegisterSubjectRoutes(r, scoreHandler)

	// 6: Swagger docs endpoint
	if os.Getenv("ENABLE_SWAGGER") == "true" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run(":" + os.Getenv("PORT"))
}