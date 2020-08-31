/*
   @Time : 20-8-27 上午11:34
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("gin server start.")
	 r:= gin.Default()
	 r.GET("/", indexHandle)

	 r.Run(":8081")
}

// index handle
func indexHandle(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"pong",
	})
}
