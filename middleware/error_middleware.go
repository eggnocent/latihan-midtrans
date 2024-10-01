package middleware

import (
	"latihanmidtrans/helper"
	"latihanmidtrans/model/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandle() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		if validationErrors(c, err) {
			return
		}

		internalServerError(c, err)
	})
}

func validationErrors(c *gin.Context, err any) bool {
	if ve, ok := err.(validator.ValidationErrors); ok {
		out := make([]web.ErrorResponse, len(ve))
		for i, fe := range ve {
			out[i] = web.ErrorResponse{
				Field:   fe.Field(),
				Message: helper.MessageForTag(fe.Tag()),
			}
		}
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   out,
		}
		c.JSON(http.StatusBadRequest, webResponse)
		c.Abort()
		return true
	}
	return false
}

func internalServerError(c *gin.Context, err any) {
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	c.JSON(http.StatusInternalServerError, webResponse)
	c.Abort()
}
