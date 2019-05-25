package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/laughmaker/go-pkg/conf"
	"github.com/laughmaker/go-pkg/file"
	"github.com/laughmaker/go-pkg/mail"
	"github.com/laughmaker/go-pkg/resp"

	"github.com/gin-gonic/gin"
)

func Try(c *gin.Context) {
	errs := recover()
	if errs == nil {
		return
	}

	fmt.Println("崩溃信息--------------------------")
	fmt.Println(errs)

	r := resp.Resp{C: c}
	r.Error(http.StatusInternalServerError)

	write2File(errs, c)

	body := formatMailRequest(c) + formatMailStack()
	mail.Send(conf.App.LogMail, fmt.Sprintf("%v", errs), body, "")

	PutLog("Error", c, fmt.Sprintf("%v", errs), string(debug.Stack()))
}

func formatMailRequest(c *gin.Context) string {
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	header, _ := json.Marshal(c.Request.Header)
	params, _ := json.Marshal(c.Params)

	body := "<strong style=\"font-size:24px;\">BASIC</strong> <br>"
	body += "method: " + c.Request.Method + "<br>"
	body += "params: " + string(params) + "<br>"
	body += "requestURI: " + c.Request.RequestURI + "<br>"
	body += "host: " + c.Request.Host + "<br>"
	body += "clientIp: " + c.ClientIP() + "<br>"
	body += "remoteAddr: " + c.Request.RemoteAddr + "<br>"

	body += "<br><strong style=\"font-size:24px;\">HEADER</strong> <br>"
	body += string(header) + "<br>"

	body += "<br><strong style=\"font-size:24px;\">BODY</strong> <br>"
	body += string(requestBody) + "<br>"

	return body
}

func formatMailStack() string {
	stack := strings.Split(string(debug.Stack()), "\n")
	str := "<br><strong style=\"font-size:24px;\">STACK</strong> <br>"
	for idx, v := range stack {
		if idx == 0 {
			str += v + "<br>"
			continue
		}

		if idx%2 == 0 {
			str += v + " </p> "
		} else {
			str += "<p style=\"padding:1px; margin:1px;\">" + v
			if idx < (len(stack) - 1) {
				str += "  <strong style=\"font-size:16px;\">-></strong> "
			}
		}

		if idx == (len(stack) - 1) {
			str += "\r\n"
		}
	}

	return str
}

func write2File(errs interface{}, c *gin.Context) {
	f, err := file.MustOpen("panic.log", conf.App.LogPanicPath)
	if err != nil {
		fmt.Println(err)
		log.Panicf("logging.Setup err:%v", err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("[%s]", time.Now().Format("2006-01-02 15:04:05")))
	f.WriteString(fmt.Sprintf("\r\n%v\r\n", errs))

	f.WriteString(formatFileRequest(c))

	f.WriteString(string(formatFileStack()))
	f.WriteString("\r\n")
}

func formatFileRequest(c *gin.Context) string {
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	header, _ := json.Marshal(c.Request.Header)
	params, _ := json.Marshal(c.Params)

	body := "method: " + c.Request.Method + "\r\n"
	body += "params: " + string(params) + "\r\n"
	body += "requestURI: " + c.Request.RequestURI + "\r\n"
	body += "host: " + c.Request.Host + "\r\n"
	body += "clientIp: " + c.ClientIP() + "\r\n"
	body += "remoteAddr: " + c.Request.RemoteAddr + "\r\n"
	body += "header:" + string(header) + "\r\n"
	body += "body:" + string(requestBody) + "\r\n"

	return body
}

func formatFileStack() string {
	stack := strings.Split(string(debug.Stack()), "\n")
	str := "[stack] \n"
	for idx, v := range stack {
		if idx == 0 {
			str += v + "\n"
			continue
		}

		if idx%2 == 0 {
			str += v + "\n "
		} else {
			str += v
			if idx < (len(stack) - 1) {
				str += " ->"
			}
		}

		if idx == (len(stack) - 1) {
			str += "\n"
		}
	}

	return str
}
