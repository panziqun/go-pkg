package resp

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Resp struct
type Resp struct {
	C *gin.Context
}

// Data response data struct
type Data struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    interface{} `json:"page"`
}

// Send resp's send data
func (r *Resp) Send(httpStatus, code int, data interface{}, page interface{}) {
	message := GetMessage(code)
	if httpStatus != http.StatusOK {
		message = http.StatusText(httpStatus)
	}

	d := Data{
		Code:    code,
		Message: message,
		Data:    data,
		Page:    page,
	}
	r.C.JSON(httpStatus, d)

	r.C.Set(RespDataKey, d)
}

// Success resp's success
func (r *Resp) Success(data interface{}) {
	r.Send(http.StatusOK, Success, data, nil)
}

// Failure resp's failure
func (r *Resp) Failure(code int) {
	r.Send(http.StatusOK, code, nil, nil)
}

// List resp's list
func (r *Resp) List(list interface{}, page interface{}) {
	r.Send(http.StatusOK, Success, list, page)
}

// Error resp's error
func (r *Resp) Error(httpStatus int) {
	r.Send(httpStatus, Failure, nil, nil)
}

// File resp's file
func (r *Resp) File(filePath string, fileName string) {
	r.C.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	r.C.Header("Content-Description", "File Transfer")
	r.C.Header("Content-Type", "application/octet-stream")
	r.C.File(filePath)
}
