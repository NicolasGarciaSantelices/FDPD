package common_function

import (
	"net/http"

	"github.com/gin-gonic/gin"

	common_models "FDPD-BACKEND/src/common/models"
)

func CreateResponse(data interface{}, lensOfData int, err error) (response common_models.Response) {

	if lensOfData == 0 {
		response = common_models.Response{
			Status:   "no content",
			Code:     204,
			Messages: "",
		}
	} else {
		if err != nil && data == nil {
			response = common_models.Response{
				Status:   "Error",
				Code:     400,
				Messages: err.Error(),
				Data:     nil,
			}
		} else {
			if data == nil {
				response = common_models.Response{
					Status:   "no content",
					Code:     204,
					Messages: "",
				}
			} else {
				response = common_models.Response{
					Status:   "OK",
					Code:     200,
					Messages: "",
					Data:     data,
				}
			}
		}
	}

	return
}

func SendResponse(c *gin.Context, response common_models.Response) {
	switch response.Code {
	case 200:
		c.JSON(http.StatusOK, response)
	case 201:
		c.JSON(http.StatusCreated, response)
	case 202:
		c.JSON(http.StatusAccepted, response)
	case 204:
		c.JSON(http.StatusNoContent, response)
	case 400:
		c.JSON(http.StatusBadRequest, response)
	case 406:
		c.JSON(http.StatusNotAcceptable, response)
	default:
		c.JSON(http.StatusInternalServerError, response)
	}
}
