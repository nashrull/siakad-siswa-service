package siswa

import (
	"context"

	"github.com/nashrull/averin/siswa/dto"
)

type SiswaUsecase interface {
	Create(ctx context.Context, payload dto.CreateDataSiswaRequest) (response dto.CreateDataSiswaResponse, err error)
	List(ctx context.Context) (response []dto.CreateDataSiswaResponse, err error)
	Update(ctx context.Context, id string, payload dto.CreateDataSiswaRequest) (response dto.CreateDataSiswaResponse, err error)
	Delete(ctx context.Context, id string) (err error)
}
