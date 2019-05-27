package address

import (
	"net/http"

	"github.com/laughmaker/go-pkg/db"
	"github.com/laughmaker/go-pkg/resp"

	"github.com/gin-gonic/gin"
)

type getAddressListForm struct {
	ProvinceCode string `form:"province_code"`
	CityCode     string `form:"city_code"`
}

// @Summary 列表
// @Tags 地址
// @Produce  json
// @Param province_code query string false "省编码"
// @Param city_code query string false "市编码"
// @Success 200 {object} resp.Data
// @Router /address/list [get]
func GetAddressList(c *gin.Context) {
	r := resp.Resp{C: c}
	var form getAddressListForm
	if err := c.ShouldBind(&form); err != nil {
		r.Error(http.StatusBadRequest)
		return
	}
	if form.CityCode != "" {
		var districts []District
		err := db.DB.Where("city_code = ?", form.CityCode).Find(&districts).Error
		if err != nil {
			panic(err)
		}
		r.Success(districts)
		return
	}
	if form.ProvinceCode != "" {
		var citys []City
		err := db.DB.Where("province_code = ?", form.ProvinceCode).Find(&citys).Error
		if err != nil {
			panic(err)
		}
		r.Success(citys)
		return
	}
	var provinces []Province
	err := db.DB.Find(&provinces).Error
	if err != nil {
		panic(err)
	}
	r.Success(provinces)
	return
}

// @Summary 国家-列表
// @Tags 地址
// @Produce  json
// @Success 200 {object} resp.Data
// @Router /address/countries [get]
func GetCountries(c *gin.Context) {
	r := resp.Resp{C: c}
	var countrys []Country
	err := db.DB.Find(&countrys).Error
	if err != nil {
		panic(err)
	}
	r.Success(countrys)
	return
}
