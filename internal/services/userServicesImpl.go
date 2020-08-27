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
	"io"
	"log"
	"time"
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

	return &pb.UserInfoResponse{Users: users}, nil
}

// 服务器流接受

func (*UserServicesImpl) GetUserInfoByServerStream(in *pb.UserInfoRequest, out pb.UserService_GetUserInfoByServerStreamServer) error {

	ret := make([]*pb.UserInfo, 0)
	var age int32 = 100
	for _, v := range in.Users {
		time.Sleep(time.Second * 2)
		user := pb.UserInfo{UserID: v.UserID,
			UserAge: age,
		}
		age++
		ret = append(ret, &user)
		if len(ret)%2 == 0 && len(ret) > 0 {
			err := out.Send(&pb.UserInfoResponse{Users: ret})
			if err != nil {
				log.Fatal(err)
			}
			ret = ret[0:0]
		}
	}
	if len(ret) > 0 {
		err := out.Send(&pb.UserInfoResponse{Users: ret})
		if err != nil {
			log.Fatal(err)
		}
	}
	ret = ret[0:0]
	return nil
}

// 客户端流
func (*UserServicesImpl) GetUserInfoByClientStream(in pb.UserService_GetUserInfoByClientStreamServer) error {
	ret := make([]*pb.UserInfo, 0)
	var err error = nil
	for {
		req, err := in.Recv()
		if err == io.EOF {// 接受完毕
			err = in.SendAndClose(&pb.UserInfoResponse{Users: ret})
			if err != nil {
				log.Fatalln(err)
			}
			return err
		}
		if err != nil {
			log.Fatalln(err)
		}

		for _, v := range req.Users {
			user := &pb.UserInfo{
				UserID:  v.UserID,
				UserAge: v.UserAge+100,
			}
			ret = append(ret, user)
		}
	}
	return err
}
