package app

import (
	"net/http"

	"github.com/laughmaker/go-pkg/conf"
	"github.com/laughmaker/go-pkg/db"
	"github.com/laughmaker/go-pkg/log"
	"github.com/laughmaker/go-pkg/mongo"
	"github.com/laughmaker/go-pkg/redis"

	"github.com/gin-gonic/gin"
)

type Application struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    interface{} `json:"page"`
}

// 加载配置组件
func Config() {
	conf.Setup()
	log.Setup()
	db.Setup()
	redis.Setup()
	mongo.Setup()
}

func (app *Application) Response(httpStatus, code int, data interface{}, page interface{}) {
	message := GetMessage(code)
	if httpStatus != http.StatusOK {
		message = http.StatusText(httpStatus)
	}

	app.C.JSON(httpStatus, Response{
		Code:    code,
		Message: message,
		Data:    data,
		Page:    page,
	})
}

func (app *Application) List(list interface{}, page Page) {
	app.Response(http.StatusOK, SUCCESS, list, page)
}

func (app *Application) Model(model interface{}) {
	app.Response(http.StatusOK, SUCCESS, model, nil)
}

func (app *Application) Failed(code int) {
	app.Response(http.StatusOK, code, nil, nil)
}

func (app *Application) Error(httpStatus int) {
	app.Response(httpStatus, FAILED, nil, nil)
}
