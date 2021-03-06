package controller

import (
	"net/http"

	"api-gin-web/utils"
	"api-gin-web/utils/errno"

	"github.com/741369/go_utils/log"
	"github.com/gin-gonic/gin"
)

// Response return model struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId interface{} `json:"trace_id"`
}

// SendResponse returns a copy of the current context that can be safely used outside the request's scope.
// This has to be used when the context has to be passed to a goroutine.
func SendResponse(context *gin.Context, err error, data interface{}) {
	acceptLanguage := context.Request.Header.Get("Accept-Language")
	language := utils.GetLanguage(acceptLanguage)
	language = "zh"
	code, message := errno.DecodeErr(err, language)
	traceId := log.GetTraceId(context)
	// always return http.StatusOK
	context.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
		TraceId: traceId,
	})
}
