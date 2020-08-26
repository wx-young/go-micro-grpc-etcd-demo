/*
   @Time : 20-8-25 下午8:59
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import (
	"context"
	pb "go-micro-grpc-etcd-demo/internal/proto"
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

	//client := pbGreet.NewGreeterClient(connet)

	//rep := &pbGreet.HelloRequest{
	//	Name: "Tom",
	//}
	//resp, err := client.SayHello(context.Background(), rep)
	//
	//if err != nil {
	//	log.Println("greet error.")
	//}
	//
	//log.Printf("greet end:  %s  ", resp.GetMessage())

	userClient := pb.NewUserServiceClient(connet)
	users := make([]*pb.UserInfo, 0)
	for i := 0; i <= 5; i++ {
		user := pb.UserInfo{
			UserID:  string(i),
			UserAge: int32(i + 1),
		}
		users = append(users, &user)
	}
	ret, err := userClient.GetUserInfo(context.Background(), &pb.UserInfoRequest{
		Users: users,
	})
	if err != nil {
		log.Fatalf("rpc get user info error: %v", err)
	}
	log.Println(ret.Users)

}
