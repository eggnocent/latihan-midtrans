package controller

import (
	"latihanmidtrans/helper"
	"latihanmidtrans/model/web"
	"latihanmidtrans/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MidtransController interface {
	Create(c *gin.Context)
}

type MidtransControllerImpl struct {
	MidtransService service.MidtransServiceImpl
}

func NewMidtransServiceImpl(midtransService service.MidtransServiceImpl) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		MidtransService: midtransService,
	}
}

func (controller *MidtransControllerImpl) Create(c *gin.Context) {
	var request web.MidtransRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		helper.PanicError(err)
	}
	midtransResponse := controller.MidtransService.Create(c, request)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   midtransResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}
