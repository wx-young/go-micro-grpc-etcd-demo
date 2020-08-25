#!/usr/bin/env bash
protoc --go_out=protobuf/greet/greet.proto=internal/proto/greet/greet.proto \
       --go_path