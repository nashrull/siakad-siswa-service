package entrypoint

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nashrull/averin/shared/config"
	handler "github.com/nashrull/averin/siswa/consumer/http"
	mysql "github.com/nashrull/averin/siswa/repository/mysql"
	"github.com/nashrull/averin/siswa/usecase"
)

func RegisterSiswaModule(
	log *log.Logger,
	e *gin.Engine,
	cfg config.Config,
) error {

	siswarepo, err := mysql.NewMysqlRepository(log, cfg)
	if err != nil {
		return err
	}

	usecasesiswa := usecase.NewSiswaUsecase(log, siswarepo)

	err = handler.NewHandler(e, usecasesiswa)
	return err
}
