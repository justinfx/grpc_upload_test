#!/bin/sh
protoc \
  --experimental_allow_proto3_optional \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  uploader/uploader.proto
  
