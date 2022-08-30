#/bin/bash

go build -v -o ./bin ./...

rm ./bin/grpc_testclient

cp ../secrets/global-entry-notify-api/* ./bin