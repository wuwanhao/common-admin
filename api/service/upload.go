// 图片上传服务层
package service

import (
	"admin-api/common/config"
	"admin-api/common/result"
	"admin-api/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"time"
)

// 定义接口和抽象方法
type IUploadService interface {
	Upload(c *gin.Context)
}

// 定义实现接口的结构体
type UploadServiceImpl struct {

}

func (u UploadServiceImpl)Upload(c *gin.Context)  {
	// 拿到文件句柄
	file, err := c.FormFile("file")
	if err != nil {
		result.Failed(c, int(result.ApiCode.FILEUPLOADERROR),
			result.ApiCode.GetMessage(result.ApiCode.FILEUPLOADERROR))
	}

	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf("%s%s%s%s",
			config.Config.ImageSettings.UploadDir,
			fmt.Sprintf("%04d", now.Year()),
			fmt.Sprintf("%02d", now.Month()),
			fmt.Sprintf("%02d", now.Day()))
	// 创建文件保存目录
	err = util.CreateDir(filePath)


	fullPath := filePath + "/" + fileName
	fmt.Println("file full path:" + fullPath)
	err = c.SaveUploadedFile(file, fullPath)
	fmt.Println(err)
	result.Success(c, fullPath)
}

var uploadService = UploadServiceImpl{}

func UploadService() IUploadService {
	return uploadService
}