#!/bin/bash

mkdir ./proto_services

protoc --go_out=./proto_services   \
    --go-grpc_out=./proto_services \
    ./doggo-proto/account.proto