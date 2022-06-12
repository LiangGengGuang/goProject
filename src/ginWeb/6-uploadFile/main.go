package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/5-models"
	"strings"
)

func upload(c *gin.Context) {
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
}

//多文件上传
func uploadMuch(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("文件获取失败", err)
		c.JSON(http.StatusBadRequest, models.ErrorResult("文件获取失败"))
		return
	}

	files := form.File["file"]
	if files == nil {
		fmt.Println("文件获取失败", err)
		c.JSON(http.StatusBadRequest, models.ErrorResult("文件获取失败"))
		return
	}

	filenames := make([]string, len(files))
	for i, file := range files {
		filename := file.Filename
		if err := c.SaveUploadedFile(file, "/Users/lianggengguang/Documents/zz_testFile/"+filename); err != nil {
			fmt.Println("文件上传失败", err)
			c.JSON(http.StatusBadRequest, models.ErrorResult("【"+filename+"】"+"文件上传失败"))
			return
		}
		filenames[i] = filename
	}

	c.JSON(http.StatusOK, models.SuccessResult("【"+strings.Join(filenames, "、")+"】"+"上传成功"))
}

func main() {

	e := gin.Default()

	e.POST("/upload", func(c *gin.Context) {
		upload(c)
	})

	//多文件上传
	e.POST("/uploadMuch", func(c *gin.Context) {
		uploadMuch(c)
	})

	e.Run(":8088")

}
