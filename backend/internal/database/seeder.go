package database

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"goldenowl-test/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

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

func SeedStudentScores(db *gorm.DB) error {
	db.Config.Logger = logger.Default.LogMode(logger.Silent)
	scores, err := loadStudentScoresFromCSV()
	if err != nil {
		return err
	}

	return seedInParallel(db, scores, 5, 1000)
}


func loadStudentScoresFromCSV() ([]models.StudentScore, error){
	// Get specific filepath
	cwd, _ := os.Getwd()
	csvPath := filepath.Join(cwd, "diem_thi_thpt_2024.csv")
	file, err := os.Open(csvPath)
	if err != nil{
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read()

	var scores []models.StudentScore
	for {
		record, err := reader.Read()
		if err == io.EOF {
            break
        }
        if err != nil || len(record) < 11 {
            continue
        }

		scores = append(scores, models.StudentScore{
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
	}



	return scores, nil

}

func seedInParallel(db *gorm.DB, data []models.StudentScore, numWorkers int, batchSize int) error {
	chunkSize := len(data) / numWorkers
	var wg sync.WaitGroup
	errChan := make(chan error, numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(data)
		}
		fmt.Println(i)
		wg.Add(1)
		go func(batch []models.StudentScore) {
			defer wg.Done()
			for i := 0; i < len(batch); i += batchSize {
				batchEnd := i + batchSize
				if batchEnd > len(batch) {
					batchEnd = len(batch)
				}
				subBatch := batch[i:batchEnd]

				err := db.Clauses(clause.OnConflict{
					DoNothing: true,
				}).CreateInBatches(subBatch, batchSize).Error

				if err != nil {
					errChan <- err
					return
				}
			}
		}(data[start:end])
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return err
	}
	return nil
}
