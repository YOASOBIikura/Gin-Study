package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func main() {
//	r := gin.Default()
//	v1 := r.Group("/v1")
//	{
//		v1.GET("/login", login)
//		v1.POST("/submit", submit)
//	}
//
//	v2 := r.Group("/v2")
//
//	{
//		v2.GET("/login", login)
//		v2.POST("/submit", submit)
//	}
//	r.Run(":8080")
//}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "刘君")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "向珂熠")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}
