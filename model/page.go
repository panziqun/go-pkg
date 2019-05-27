package model

import (
	"context"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/laughmaker/go-pkg/mongo"
)

// Page page data struct
type Page struct {
	PageIndex int64 `json:"page_index"`  // 分页索引
	PageSize  int64 `json:"page_size"`   // 分页大小
	PageCount int64 `json:"total_pages"` // 分页数量
	TotalRows int64 `json:"total_rows"`  // 总行数
}

// DbPage get sqldb page
func DbPage(c *gin.Context, db *gorm.DB) Page {
	p := getPage(c)

	db.Count(&p.TotalRows)
	p.PageCount = int64(math.Ceil(float64(p.TotalRows) / float64(p.PageSize)))
	return p
}

// MongoPage get mongodb page
func MongoPage(c *gin.Context, dbname string, name string, filter interface{}) Page {
	p := getPage(c)

	m := mongo.Mongo{DB: mongo.GetDB(dbname)}
	p.TotalRows, _ = m.DB.Collection(name).CountDocuments(context.Background(), filter)
	p.PageCount = int64(math.Ceil(float64(p.TotalRows) / float64(p.PageSize)))
	return p
}

func getPage(c *gin.Context) Page {
	p := Page{}
	p.PageIndex, _ = strconv.ParseInt(c.DefaultQuery("page_index", "1"), 10, 64)
	if p.PageIndex <= 0 {
		p.PageIndex = 1
	}
	p.PageSize, _ = strconv.ParseInt(c.DefaultQuery("page_size", "20"), 10, 64)
	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = 20
	}

	return p
}

// Limit get page limit
func (p *Page) Limit() int64 {
	return p.PageSize
}

// Offset get page offset
func (p *Page) Offset() int64 {
	return (p.PageIndex - 1) * p.Limit()
}
