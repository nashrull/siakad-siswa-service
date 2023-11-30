package usecase

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/nashrull/averin/siswa"
	domain "github.com/nashrull/averin/siswa"
	"github.com/nashrull/averin/siswa/dto"
)

type Usecase struct {
	log  *log.Logger
	repo domain.SiswaRepository
}

// Delete implements siswa.SiswaUsecase.
func (s *Usecase) Delete(ctx context.Context, id string) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// List implements siswa.SiswaUsecase.
func (r *Usecase) List(ctx context.Context) (response []dto.CreateDataSiswaResponse, err error) {
	response = make([]dto.CreateDataSiswaResponse, 0)

	result, err := r.repo.List(ctx, domain.Siswa{})
	if err != nil {
		return response, err
	}

	for _, v := range result {
		var t dto.CreateDataSiswaResponse
		t.ID = strconv.Itoa(int(v.ID))
		t.Alamat = v.Alamat
		t.Name = v.Nama
		t.Alamat = v.Alamat
		t.TempatLahir = v.TempatLahir
		t.TanggalLahir = v.TanggalLahir.String()
		response = append(response, t)
	}

	return response, nil
}

// Update implements siswa.SiswaUsecase.
func (u *Usecase) Update(ctx context.Context, id string, payload dto.CreateDataSiswaRequest) (response dto.CreateDataSiswaResponse, err error) {
	tanggalLahir, _ := time.Parse("2006-01-02", payload.TanggalLahir)
	edited := siswa.Siswa{
		Nama:         payload.Name,
		Alamat:       payload.Alamat,
		TempatLahir:  payload.TempatLahir,
		TanggalLahir: tanggalLahir,
	}

	err = u.repo.Save(ctx, id, edited)
	if err != nil {
		return response, err
	}

	response = dto.CreateDataSiswaResponse{
		ID:                     id,
		CreateDataSiswaRequest: payload,
	}

	return response, nil

}

// Create implements siswa.SiswaUsecase.
func (u *Usecase) Create(ctx context.Context, payload dto.CreateDataSiswaRequest) (response dto.CreateDataSiswaResponse, err error) {
	if !payload.AgeIsLessSeven() {
		return response, errors.New("usia kurang dari 7")
	}

	var obj siswa.Siswa
	obj.Alamat = payload.Alamat
	obj.Nama = payload.Name
	obj.TempatLahir = payload.TempatLahir
	obj.TanggalLahir, _ = time.Parse("2006-01-02", payload.TanggalLahir)
	u.log.Println(obj)
	insertedID, err := u.repo.Create(ctx, obj)
	if err != nil {
		return response, err
	}

	response.ID = strconv.Itoa(int(insertedID))
	response.CreateDataSiswaRequest = payload

	return response, nil
}

func NewSiswaUsecase(log *log.Logger, repo domain.SiswaRepository) domain.SiswaUsecase {
	return &Usecase{
		log:  log,
		repo: repo,
	}
}
