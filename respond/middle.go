package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//func main() {
//	r := gin.Default()
//	r.Use(MiddleWare())
//	r.GET("/ce", func(c *gin.Context) {
//		req, _ := c.Get("name")
//		fmt.Println("name", req)
//		c.JSON(http.StatusOK, gin.H{"name": req})
//	})
//	r.Run(":8080")
//}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		c.Set("name", "刘君")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
