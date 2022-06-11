package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//中间件
func nextMiddleware(c *gin.Context) {
	fmt.Println("====>执行nextMiddleware成功<====")
	start := time.Now().UnixNano()

	//获取中间件的共享值
	if value, ok := c.Get("key"); ok == true {
		fmt.Printf("value:%s\n", value)
	}

	//执行剩余的处理程序
	c.Next()
	end := time.Now().UnixNano()
	fmt.Println("====>执行nextMiddleware完毕<====")

	fmt.Printf("执行时间：%d秒", (end-start)/1e9)

}

func abortMiddleware(c *gin.Context) {
	fmt.Println("====>执行abortMiddleware成功<====")
	start := time.Now().UnixNano()

	//只执行当前剩余的处理程序
	c.Abort()
	end := time.Now().UnixNano()
	fmt.Println("====>执行abortMiddleware完毕<====")

	fmt.Printf("执行时间：%d秒", (end-start)/1e9)

}

func groupMiddleware(c *gin.Context) {
	fmt.Println("====>执行groupMiddleware成功<====")

	//配置中间件的共享值
	c.Set("key", "nextMiddleware")

	time.Sleep(time.Second)
	fmt.Println("====>执行groupMiddleware完毕<====")
}

//中间件中使用协程
func goroutineMiddleware(c *gin.Context) {

	fmt.Println("====>执行goroutineMiddleware成功<====")

	//通过c.Copy()复制出一个Context
	cCopy := c.Copy()
	go func() {
		fmt.Printf("goroutine协程调用，path：%s\n", cCopy.Request.URL.Path)
	}()

	fmt.Println("====>执行goroutineMiddleware完毕<====")
}

func main() {

	e := gin.Default()

	//全局中间件(配置于路由组的中间件后面，会失效)
	//e.Use(abortMiddleware)

	group := e.Group("/get")
	group.Use(groupMiddleware, goroutineMiddleware)
	{
		group.GET("/one", nextMiddleware, func(c *gin.Context) {
			time.Sleep(2 * time.Second)
			fmt.Println("====>one:返回请求结果<====")
			c.JSON(http.StatusOK, gin.H{
				"msg": "success",
			})
		})

		group.GET("/two", nextMiddleware, func(c *gin.Context) {
			time.Sleep(2 * time.Second)
			fmt.Println("====>two:返回请求结果<====")
			c.JSON(http.StatusOK, gin.H{
				"msg": "success",
			})
		})
	}

	e.Run(":8088")

}
