/*
   @Time : 20-8-25 下午3:28
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import (
	"log"
	"time"
)

func main() {
	log.Println("grpc demo.")
	a := make(chan int, 1)
	b := make(chan int, 1)
	c := make(chan int, 1)
	go func() {
		a <- 1
		b <- 1
		c <- 1
	}()

	time.Sleep(3*time.Second)
	select {
	case <-a:
		log.Println("a")
	case <-b:
		log.Println("b")
	case <-c:
		log.Println("c")
	}
}
