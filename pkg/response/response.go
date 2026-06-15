package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	c.JSON(200, Response{
		Code:    0,
		Message: "Success",
		Data:    data,
	})
}

func Fail(c *gin.Context, code int, message string) {
	c.JSON(400, Response{
		Code:    code,
		Message: message,
	})
}
