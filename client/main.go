/*
   @Time : 20-8-25 下午8:59
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import (
	"context"
	pbGreet "go-micro-grpc-etcd-demo/internal/proto/greet"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.Println("start grpc_client")

	connet, err := grpc.Dial("localhost:8034", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc dial error.")
	}
	defer connet.Close()

	client := pbGreet.NewGreeterClient(connet)

	rep := &pbGreet.HelloRequest{
		Name:"Tom",
	}
	resp, err := client.SayHello(context.Background(), rep)

	if err != nil {
		log.Println("greet error.")
	}

	log.Printf("greet end:  %s  ",resp.GetMessage())

}
