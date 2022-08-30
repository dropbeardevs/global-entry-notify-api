#/bin/bash

mkdir ./bin

go build -v -o ./bin ./...

rm ./bin/grpc_testclient

cp ../secrets/global-entry-notify-api/* ./bin