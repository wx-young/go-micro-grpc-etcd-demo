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
	"time"
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
	ctx := context.Background()

	//users := make([]*pb.UserInfo, 0)
	//for i := 0; i <= 15; i++ {
	//	user := pb.UserInfo{
	//		UserID:  string(i),
	//		UserAge: int32(i + 1),
	//	}
	//	users = append(users, &user)
	//}
	// 普通接受
	//ret, err := userClient.GetUserInfo(context.Background(), &pb.UserInfoRequest{
	//	Users: users,
	//})
	//if err != nil {
	//	log.Fatalf("rpc get user info error: %v", err)
	//}
	//log.Println(ret.Users)

	// 服务器流接受
	//sream, err := userClient.GetUserInfoByServerStream(context.Background(), &pb.UserInfoRequest{
	//	Users: users,
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for {
	//	ret, err := sream.Recv()
	//	if err == io.EOF {
	//		// 完成
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	log.Println(ret)
	//}

	// 客户端流发送
	stream, err := userClient.GetUserInfoByClientStream(ctx)
	if err != nil{
		log.Fatalln(err)
	}
	users := make([]*pb.UserInfo, 0)

	for j := 0; j < 14; j++ {
		time.Sleep(time.Second)
		user := &pb.UserInfo{
			UserID:  string(1000 + j),
			UserAge: int32(j + 20),
		}
		users = append(users, user)
		if len(users)%2 == 0 && len(users)>0 {
			err := stream.Send(&pb.UserInfoRequest{Users:users})
			log.Println(users)
			if err != nil {
				log.Fatalln(err)
			}
			users= users[0:0]
		}
	}
	if len(users) >0 {
		err := stream.Send(&pb.UserInfoRequest{Users:users})
		log.Println(users)
		if err != nil {
			log.Fatalln(err)
		}
		users = users[0:0]
	}

	//  关闭
	ret , err:= stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(ret.Users)
}
