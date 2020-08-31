/*
   @Time : 20-8-27 下午10:51
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"log"
)

func main() {
	log.Println("go-micro server start.")

	ginRouter := gin.Default()
	ginRouter.GET("/user", getUser)

	consulRegister := consul.NewRegistry()
	server := web.NewService(
		web.Name("produsService"),
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry())
}

func getUser(c *gin.Context){
	c.String(200, "user api.")
}