package services

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"goldenowl-test/internal/models"
	"goldenowl-test/internal/repositories"
)

type StudentScoreService struct {
    repo *repositories.StudentScoreRepo
}

func NewStudentScoreService(repo *repositories.StudentScoreRepo) *StudentScoreService {
    return &StudentScoreService{repo: repo}
}

func (s *StudentScoreService) GetFirst10Rows() (*[]models.StudentScore, error) {
    return s.repo.GetFirst10Rows()
}

func (s *StudentScoreService) GetScoreBySBD(sbd string) (*models.StudentScore, error){
	return s.repo.GetScoreBySBD(sbd)
}

func (s *StudentScoreService) GetScoreReportBySubject() (map[string]map[string]int64, error){
	return s.repo.GetScoreReportBySubject()
}

func (s *StudentScoreService) GetTop10GroupA() (*[]models.StudentScore, error){
	return s.repo.GetTop10GroupA()
}

func parseFloat(value string) *float64 {
    if value == "" {
        return nil
    }
    f, err := strconv.ParseFloat(value, 64)
    if err != nil {
        return nil
    }
    return &f
}

func (s *StudentScoreService) SeedStudentScores() error {
    count, err := s.repo.Count()
    if err != nil {
        return err
    }
    if count > 0 {
        log.Println("Student scores already seeded.")
        return nil
    }

    cwd, _ := os.Getwd()
    csvPath := filepath.Join(cwd, "diem_thi_thpt_2024.csv")
    file, err := os.Open(csvPath)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    // skip header
    if _, err := reader.Read(); err != nil {
        return err
    }

    var batch []models.StudentScore
    countBatchSize := 0
    var copyRowsCount int64 = 0
    batchSize := 10000

    for {
        record, err := reader.Read()
        if err != nil {
            if err.Error() == "EOF" {
                break
            }
            return err
        }
        if len(record) < 11 {
            continue
        }

        batch = append(batch, models.StudentScore{
            SBD:        record[0],
            Toan:       parseFloat(record[1]),
            NguVan:     parseFloat(record[2]),
            NgoaiNgu:   parseFloat(record[3]),
            VatLi:      parseFloat(record[4]),
            HoaHoc:     parseFloat(record[5]),
            SinhHoc:    parseFloat(record[6]),
            LichSu:     parseFloat(record[7]),
            DiaLi:      parseFloat(record[8]),
            GDCD:       parseFloat(record[9]),
            MaNgoaiNgu: record[10],
        })
        countBatchSize++

        if countBatchSize == batchSize {
            copyRows, err := s.repo.CopyBatch(batch)
            if err != nil {
                return err
            }
            copyRowsCount += copyRows
            batch = nil // reset to free up
            countBatchSize = 0
        }
    }

    if len(batch) > 0 {
        copyRows, err := s.repo.CopyBatch(batch)
        if err != nil {
            return err
        }
        copyRowsCount += copyRows
    }

    log.Printf("Seeded %d student scores from CSV\n", copyRowsCount)
    return nil
}
