#/bin/sh

rm -rf ./bin

mkdir ./bin

go build -v -o ./bin ./...

rm ./bin/grpc_testclient

cp ../secrets/global-entry-notify-api/global-entry-c8373-fe72e1ae9c11.json ./bin/fbauth.json
cp ../secrets/global-entry-notify-api/config.prod.yml ./bin/config.yml

docker build -t global-entry-notify-api .