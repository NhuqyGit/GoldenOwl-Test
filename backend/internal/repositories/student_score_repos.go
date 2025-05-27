package repositories

import (
	"goldenowl-test/internal/database"

	"gorm.io/gorm"
)

type StudentScoreRepo struct {
	db *gorm.DB
}

func NewStudentScoreRepo() *StudentScoreRepo{
	return &StudentScoreRepo{db: database.GetDB()}
}




