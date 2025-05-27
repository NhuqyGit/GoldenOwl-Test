package repositories

import (
	"goldenowl-test/internal/database"
	"goldenowl-test/internal/models"

	"gorm.io/gorm"
)

type StudentScoreRepo struct {
	db *gorm.DB
}

func NewStudentScoreRepo() *StudentScoreRepo{
	return &StudentScoreRepo{db: database.GetDB()}
}

func (r *StudentScoreRepo) GetFirst10Rows() (*[]models.StudentScore, error) {
	var scores []models.StudentScore
	result := r.db.Limit(10).Find(&scores)
	if result.Error != nil {
		return nil, result.Error
	}
	return &scores, nil
}


