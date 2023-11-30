package siswa

import (
	"context"
)

type SiswaRepository interface {
	Create(ctx context.Context, payload Siswa) (insertedID uint, err error)
	List(ctx context.Context, payload Siswa) (response []Siswa, err error)
	Save(ctx context.Context, id string, payload Siswa) (err error)
	Delete(ctx context.Context, id string) (err error)
}
