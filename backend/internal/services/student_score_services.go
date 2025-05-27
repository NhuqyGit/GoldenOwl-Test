package services

import (
	"goldenowl-test/internal/models"
	"goldenowl-test/internal/repositories"
)

type StudentScoreService struct {
	repo *repositories.StudentScoreRepo
}

func NewStudentScoreService(repo *repositories.StudentScoreRepo) *StudentScoreService{
	return &StudentScoreService{repo: repo}
}

func (s *StudentScoreService) GetFirst10Rows() (*[]models.StudentScore, error) {
	return s.repo.GetFirst10Rows()
}