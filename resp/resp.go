package resp

import (
	"app/conf"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	C *gin.Context
}

type Data struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    interface{} `json:"page"`
}

func (r *Resp) Send(httpStatus, code int, data interface{}, page interface{}) {
	message := conf.GetMessage(code)
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

	r.C.Set(conf.RespDataKey, d)
}

func (r *Resp) Success(data interface{}) {
	r.Send(http.StatusOK, conf.Success, data, nil)
}

func (r *Resp) Failure(code int) {
	r.Send(http.StatusOK, code, nil, nil)
}

func (r *Resp) List(list interface{}, page interface{}) {
	r.Send(http.StatusOK, conf.Success, list, page)
}

func (r *Resp) Error(httpStatus int) {
	r.Send(httpStatus, conf.Failure, nil, nil)
}

func (r *Resp) File(filePath string, fileName string) {
	r.C.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(fileName))
	r.C.Header("Content-Description", "File Transfer")
	r.C.Header("Content-Type", "application/octet-stream")
	r.C.File(filePath)
}
