/*
   @Time : 20-8-26 下午6:12
   @Author : young
   @File : userServicesImp
   @Software: GoLand
*/
package services

import (
	"context"
	pb "go-micro-grpc-etcd-demo/internal/proto"
	"log"
)

type UserServicesImpl struct {
}

func (t *UserServicesImpl) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	users := make([]*pb.UserInfo, 0)
	var uid = "100"
	var age int32 = 10
	for _, v := range in.Users {
		user := pb.UserInfo{}
		user.UserID = uid
		user.UserAge = age
		age++
		log.Println(v)
		//user.UserID =
		users = append(users, &user)
	}

	return &pb.UserInfoResponse{Users:users}, nil
}
