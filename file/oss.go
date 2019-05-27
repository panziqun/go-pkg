package file

import (
	"fmt"
	"mime/multipart"
	"path"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/laughmaker/go-pkg/conf"
	"github.com/laughmaker/go-pkg/resp"
	"github.com/laughmaker/go-pkg/util"
)

// @Summary 上传
// @Tags 文件
// @Produce  json
// @Success 200 {object} resp.Data
// @Router /file/upload [post]
func UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}

	url := put2OSS(fileHeader)
	r := resp.Resp{C: c}
	r.Success(url)
}

func put2OSS(fh *multipart.FileHeader) string {
	ossConf := conf.Section("aliyunOSS")
	client, err := oss.New(ossConf["Endpoint"], ossConf["AccessKeyId"], ossConf["AccessKeySecret"])
	if err != nil {
		panic(err)
	}

	bucket, err := client.Bucket(ossConf["Bucket"])
	if err != nil {
		panic(err)
	}

	file, err := fh.Open()
	if err != nil {
		panic(err)
	}

	ext := path.Ext(fh.Filename)
	filename := fmt.Sprintf("%d%s%s", time.Now().Unix(), util.Md5(fh.Filename), ext)
	err = bucket.PutObject(filename, file)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s%s", ossConf["Host"], filename)
}
