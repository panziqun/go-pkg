package kdniao

import (
	"encoding/json"
	"net/url"

	"github.com/laughmaker/go-pkg/conf"
	"github.com/laughmaker/go-pkg/curl"
	"github.com/laughmaker/go-pkg/db"
	"github.com/laughmaker/go-pkg/resp"
	"github.com/laughmaker/go-pkg/util"

	"github.com/gin-gonic/gin"
)

const RealTimeLogisticUrl = "http://api.kdniao.com/Ebusiness/EbusinessOrderHandle.aspx"

func KdniaoTrack() (result Response, err error) {
	var shipperCode, logisticCode string = "STO", "6610651980589"
	kdConf := conf.Section("kdniao")
	requestData, _ := json.Marshal(map[string]string{
		"ShipperCode":  shipperCode,
		"LogisticCode": logisticCode,
	})
	body := map[string]interface{}{
		"RequestData": url.QueryEscape(string(requestData)),
		"EBusinessID": kdConf["MchId"],
		"RequestType": "8001",
		"DataSign":    string(util.EncodeBase64([]byte(util.Md5(string(requestData) + kdConf["ApiKey"])))),
		"DataType":    "2",
	}
	response, err := curl.NewRequest().SetUrl(RealTimeLogisticUrl).SetBody(body).Form()
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(response.Body), &result)
	if err != nil {
		return result, err
	}
	return
}

// @Summary 物流公司-列表
// @Tags 快递鸟
// @Produce  json
// @Success 200 {object} resp.Data
// @Router /kdniao/expresses [get]
func GetExpresses(c *gin.Context) {
	r := resp.Resp{C: c}
	var express []Express
	if err := db.DB.Order("sort asc").Find(&express).Error; err != nil {
		panic(err)
	}
	r.Success(express)
}
