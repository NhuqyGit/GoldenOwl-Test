package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBProvider interface {
	GetDB() *gorm.DB
	GetPgxPool() *pgxpool.Pool
}

type Database struct {
	db      *gorm.DB
	pgxPool *pgxpool.Pool
}

var (
	instance *Database
	once     sync.Once
)

// NewDatabase singleton
func NewDatabase() DBProvider {
	once.Do(func() {
		instance = &Database{}
		if err := instance.init(); err != nil {
			log.Fatal("Failed to initialize database:", err)
		}
	})
	return instance
}

func (d *Database) init() error {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSL")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSSL)

	var err error
	d.db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gorm connection failed: %w", err)
	}

	d.pgxPool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		return fmt.Errorf("pgxpool connection failed: %w", err)
	}

	log.Println("Connected to Database! (GORM + pgxpool)")

	Migrate(d.db)
	return nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

func (d *Database) GetPgxPool() *pgxpool.Pool {
	return d.pgxPool
}
