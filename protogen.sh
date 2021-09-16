#!/bin/bash

protoc -I ./proto/product product.proto \
  --js_out=import_style=commonjs:./client/src \
  --grpc-web_out=import_style=commonjs,mode=grpcweb:./client/src


protoc -I ./proto/product product.proto \
  --go_out=./proto/product --go_opt=paths=source_relative \
  --go-grpc_out=./proto/product --go-grpc_opt=paths=source_relative
