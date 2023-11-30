package http

import (
	"github.com/gin-gonic/gin"

	httpresponse "github.com/nashrull/averin/shared/http"
	"github.com/nashrull/averin/siswa"
	"github.com/nashrull/averin/siswa/dto"
)

type Handler struct {
	usecase siswa.SiswaUsecase
}

func NewHandler(e *gin.Engine, u siswa.SiswaUsecase) error {
	handler := &Handler{
		usecase: u,
	}

	siswaRoute := e.Group("/siswa")
	siswaRoute.GET("/", handler.GetAll)
	siswaRoute.POST("/", handler.Create)
	siswaRoute.DELETE("/:id", handler.Delete)
	siswaRoute.PUT("/:id", handler.Edit)
	return nil
}

func (h *Handler) GetAll(c *gin.Context) {
	response, err := h.usecase.List(c)
	if err != nil {
		httpresponse.Response(c, err, nil)
		return
	}

	httpresponse.Response(c, nil, response)
}

func (h *Handler) Create(c *gin.Context) {
	var payload dto.CreateDataSiswaRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		httpresponse.Response(c, err, nil)
		return
	}
	response, err := h.usecase.Create(c, payload)
	if err != nil {
		httpresponse.Response(c, err, nil)
		return
	}

	httpresponse.Response(c, nil, response)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.usecase.Delete(c, id)
	if err != nil {
		httpresponse.Response(c, err, nil)
		return
	}

	httpresponse.Response(c, nil, nil)
}

func (h *Handler) Edit(c *gin.Context) {
	id := c.Param("id")
	var payload dto.CreateDataSiswaRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		httpresponse.Response(c, err, nil)
		return
	}

	response, err := h.usecase.Update(c, id, payload)
	if err != nil {
		httpresponse.Response(c, err, nil)
		return
	}

	httpresponse.Response(c, nil, response)
}
