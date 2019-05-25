package es

import (
	"context"
	"fmt"
	"go-pkg/conf"
	"log"
	"os"
	"strconv"

	"github.com/olivere/elastic/v7"
)

// Client elasticsearch client
var Client *elastic.Client

// Setup config elasticsearch
func Setup() {
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	section := conf.Section("elasticsearch")
	Client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(section["Host"]), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	info, code, err := Client.Ping(section["Host"]).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

// IndexExists 校验 index 是否存在
func IndexExists(index ...string) bool {
	exists, err := Client.IndexExists(index...).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return exists
}

// CreateIndex 创建 index
func CreateIndex(index, mapping string) bool {
	result, err := Client.CreateIndex(index).BodyString(mapping).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return result.Acknowledged
}

// DelIndex 删除 index
func DelIndex(index ...string) bool {
	response, err := Client.DeleteIndex(index...).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return response.Acknowledged
}

// Add 新增文档
func Add(index string, data interface{}) string {
	response, err := Client.Index().Index(index).BodyJson(data).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return response.Id
}

// BulkAdd 批量新增
func BulkAdd(index string, datas ...interface{}) {
	bulkRequest := Client.Bulk()
	for i, data := range datas {
		doc := elastic.NewBulkIndexRequest().Index(index).Id(strconv.Itoa(i)).Doc(data)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(context.TODO())
	if err != nil {
		panic(err)
	}
	failed := response.Failed()
	iter := len(failed)
	fmt.Printf("error: %v, %v\n", response.Errors, iter)
}

// GetDoc 获取指定 Id 的文档
func GetDoc(index, id string) []byte {
	response, err := Client.Get().Index(index).Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	source, err := response.Source.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return source
}
