package dto

import (
	"time"
)

type CreateDataSiswaRequest struct {
	Name         string `json:"nama" binding:"required"`
	Alamat       string `json:"alamat" binding:"required"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir" binding:"required"`
}

func (s CreateDataSiswaRequest) AgeIsLessSeven() bool {
	parsedTime, _ := time.Parse("2006-01-02", s.TanggalLahir)

	// semisal siswa harus berusia lebih dari 1 tahun di sekolah
	today := time.Now()
	umur := today.Year() - parsedTime.Year()
	if today.Month() < parsedTime.Month() ||
		(today.Month() == parsedTime.Month() && today.Day() < parsedTime.Day()) {
		umur--
	}
	if umur < 7 {
		return true
	}
	return false
}

type CreateDataSiswaResponse struct {
	ID string `json:"id" binding:"required"`
	CreateDataSiswaRequest
}
