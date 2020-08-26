/*
   @Time : 20-8-26 上午9:52
   @Author : young
   @File : greetServices
   @Software: GoLand
*/
package services

import (
	"context"
	"errors"
)
import pbGreet "go-micro-grpc-etcd-demo/internal/proto/greet"
//import "google.golang.org/grpc/codes"
//import "google.golang.org/grpc/status"

type GreetServiceImpl struct {
}

func (*GreetServiceImpl) SayHello(ctx context.Context, req *pbGreet.HelloRequest) (*pbGreet.HelloReply, error) {
	if ctx.Err() == context.Canceled {
		//return nil, status.Errorf(codes.Canceled, "searchService.Search canceled")
		return nil, errors.New("nothing")
	}
	reply := pbGreet.HelloReply{
		Message: "Hello  " + req.Name,
	}
	return &reply, nil
}
