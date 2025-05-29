package repositories

import (
	"context"
	"fmt"
	"goldenowl-test/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"
)

type StudentScoreRepo struct {
    db      *gorm.DB
    pgxPool *pgxpool.Pool
}

func NewStudentScoreRepo(db *gorm.DB, pgxPool *pgxpool.Pool) *StudentScoreRepo {
    return &StudentScoreRepo{db: db, pgxPool: pgxPool}
}

func (r *StudentScoreRepo) GetFirst10Rows() (*[]models.StudentScore, error) {
    var scores []models.StudentScore
    if err := r.db.Limit(10).Find(&scores).Error; err != nil {
        return nil, err
    }
    return &scores, nil
}

func (r *StudentScoreRepo) GetScoreBySBD(sbd string) (*models.StudentScore, error) {
    var score models.StudentScore
    if err := r.db.Where("sbd = ?", sbd).First(&score).Error; err != nil {
        return nil, err
    }
    return &score, nil
}

func (r *StudentScoreRepo) GetScoreReportBySubject() (map[string]map[string]int64, error) {
	subjects := []string{"toan", "ngu_van", "ngoai_ngu", "vat_li", "hoa_hoc", "sinh_hoc", "lich_su", "dia_li", "gdcd"}
	ranges := []struct {
		Label string
		Min   float64
		Max   float64
	}{
		{"Poor", 0, 4},
		{"Average", 4, 6},
		{"Good", 6, 8},
		{"Excellent", 8, 11},
	}

	result := make(map[string]map[string]int64)

	for _, subject := range subjects {
		result[subject] = make(map[string]int64)
		for _, rg := range ranges {
			var count int64
			query := r.db.Model(&models.StudentScore{}).
				Where(fmt.Sprintf("%s >= ? AND %s < ?", subject, subject), rg.Min, rg.Max).
				Count(&count)
			if query.Error != nil {
				return nil, query.Error
			}
			result[subject][rg.Label] = count
		}
	}

	return result, nil
}

func (r *StudentScoreRepo) GetTop10GroupA() (*[]models.StudentScore, error) {
	var scores []models.StudentScore

	err := r.db.
		Model(&models.StudentScore{}).
		Select("*").
		Order("COALESCE(toan, 0) + COALESCE(vat_li, 0) + COALESCE(hoa_hoc, 0) DESC, id ASC").
		Limit(10).
		Find(&scores).Error

	if err != nil {
		return nil, err
	}
	return &scores, nil
}


func (r *StudentScoreRepo) Count() (int64, error) {
    var count int64
    if err := r.db.Model(&models.StudentScore{}).Count(&count).Error; err != nil {
        return 0, err
    }
    return count, nil
}

func (r *StudentScoreRepo) CopyBatch(batch []models.StudentScore) (int64, error) {
    columns := []string{
        "sbd", "toan", "ngu_van", "ngoai_ngu", "vat_li",
        "hoa_hoc", "sinh_hoc", "lich_su", "dia_li", "gdcd", "ma_ngoai_ngu",
    }

    copyCount, err := r.pgxPool.CopyFrom(
        context.Background(),
        pgx.Identifier{"student_scores"},
        columns,
        pgx.CopyFromSlice(len(batch), func(i int) ([]any, error) {
            r := batch[i]
            return []any{
                r.SBD, r.Toan, r.NguVan, r.NgoaiNgu, r.VatLi,
                r.HoaHoc, r.SinhHoc, r.LichSu, r.DiaLi, r.GDCD, r.MaNgoaiNgu,
            }, nil
        }),
    )
    return copyCount, err
}
