package httpresponse

import "github.com/gin-gonic/gin"

type WebResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(500, WebResponse{
			Code:    500,
			Status:  false,
			Message: err.Error(),
			Data:    make([]map[string]interface{}, 0),
		})
		return
	} else {
		if data == nil {
			data = gin.H{}
		}
		c.JSON(200, WebResponse{
			Code:    200,
			Status:  true,
			Message: "berhasil",
			Data:    data,
		})
		return
	}

}
