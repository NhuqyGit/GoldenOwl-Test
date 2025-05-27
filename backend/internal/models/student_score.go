package models

type StudentScore struct {
	ID         uint     `gorm:"primaryKey"`
	SBD        string   `csv:"sbd" gorm:"uniqueIndex"`
	Toan       *float64 `csv:"toan"`
	NguVan     *float64 `csv:"ngu_van"`
	NgoaiNgu   *float64 `csv:"ngoai_ngu"`
	VatLi      *float64 `csv:"vat_li"`
	HoaHoc     *float64 `csv:"hoa_hoc"`
	SinhHoc    *float64 `csv:"sinh_hoc"`
	LichSu     *float64 `csv:"lich_su"`
	DiaLi      *float64 `csv:"dia_li"`
	GDCD       *float64 `csv:"gdcd"`
	MaNgoaiNgu string   `csv:"ma_ngoai_ngu"`
}
