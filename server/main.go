/*
   @Time : 20-8-25 下午8:59
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import (
	pb "go-micro-grpc-etcd-demo/internal/proto"
	"go-micro-grpc-etcd-demo/internal/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

const
(
	HOST string = "localhost"
	PORT string = "8034"
)

func main() {
	log.Printf("server start.")
	listen, err := net.Listen("tcp", ":8034")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := listen.Close()
		if err != nil {
			log.Println("listen close error.")
		}
	}()

	//rpcServer := grpc.NewServer()
	//pbGreet.RegisterGreeterServer(rpcServer, new(services.GreetServiceImpl))
	//err = rpcServer.Serve(listen)
	//if err != nil {
	//	log.Fatalln("greet service start fail")
	//}

	rpcUser := grpc.NewServer()
	pb.RegisterUserServiceServer(rpcUser, new(services.UserServicesImpl))
	err = rpcUser.Serve(listen)
	if err != nil {
		log.Fatalf("user service start fail.")
	}

}
