package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/5-models"
)

func main() {

	e := gin.Default()

	e.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("文件获取失败", err)
			c.JSON(http.StatusBadRequest, models.ErrorResult("文件获取失败"))

			return
		}

		filename := file.Filename
		if err := c.SaveUploadedFile(file, "/Users/lianggengguang/Documents/zz_testFile/"+filename); err != nil {
			fmt.Println("文件上传失败", err)
			c.JSON(http.StatusBadRequest, models.ErrorResult(filename+"文件上传失败"))
			return
		}

		c.JSON(http.StatusOK, models.SuccessResult(filename+"上传成功"))
	})

	e.Run(":8088")

}
