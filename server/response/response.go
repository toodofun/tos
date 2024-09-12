package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(ctx *gin.Context, httpCode int, code string, message string, data interface{}) {
	ctx.JSON(httpCode, &Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Success(ctx *gin.Context, data any) {
	New(ctx, http.StatusOK, CodeSuccess, "success", data)
}

func Error(ctx *gin.Context, code string) {
	New(ctx, http.StatusOK, code, message(code), nil)
}

func ErrorWithMsg(ctx *gin.Context, code string, message string) {
	New(ctx, http.StatusOK, code, message, nil)
}
