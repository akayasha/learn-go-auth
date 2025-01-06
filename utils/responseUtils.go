package utils

import (
	"github.com/gin-gonic/gin"
)

type StandardResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Respond(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, StandardResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func RespondError(c *gin.Context, code int, message string) {
	c.JSON(code, StandardResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
