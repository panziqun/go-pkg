package kdniao

import (
	"github.com/laughmaker/go-pkg/model"
)

type Response struct {
	EBusinessID  string      `json:"EBusinessID"`
	OrderCode    string      `json:"OrderCode"`
	ShipperCode  string      `json:"ShipperCode"`
	LogisticCode string      `json:"LogisticCode"`
	Success      bool        `json:"Success"`
	State        int         `json:"State"`
	Reason       interface{} `json:"Reason"`
	Traces       []Trace     `json:"Traces"`
}

type Trace struct {
	AcceptTime    string      `json:"AcceptTime"`
	AcceptStation string      `json:"AcceptStation"`
	Remark        interface{} `json:"Remark"`
}

type Express struct {
	Name string `json:"name" gorm:"name"`
	Code string `json:"code" gorm:"code"`
	Sort int    `json:"sort" gorm:"sort"`
	Type string `json:"type" gorm:"type"`
	model.Model
}
