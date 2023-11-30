package siswa

import (
	"time"

	"github.com/google/uuid"
)

type Siswa struct {
	ID           uint
	Nama         string `gorm:"column:nama"`
	Alamat       string
	TempatLahir  string
	TanggalLahir time.Time `gorm:"column:tanggal_lahir;type:date"`
}

func (s Siswa) CreateUUID() uuid.UUID {
	return uuid.New()
}
