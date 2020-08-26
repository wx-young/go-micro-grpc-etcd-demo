#!/usr/bin/env bash
# 没有rpc的情况下
# protoc --go_out=. ./protobuf/greet/greet.proto
protoc  --go_out=plugins=grpc:. ./protobuf/greet/greet.proto

# models
protoc   --go_out=plugins=grpc:. ./protobuf/models.proto

# user
protoc  --go_out=plugins=grpc:. ./protobuf/user.proto