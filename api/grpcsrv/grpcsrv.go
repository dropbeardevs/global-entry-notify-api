package grpcsrv

import (
	"net"
	"sync"

	pb "bitbucket.org/dropbeardevs/global-entry-notify-api/api/proto"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"google.golang.org/grpc"
)

type ServiceServer struct {
	pb.UnimplementedGlobalEntryNotifyServiceServer
}

var serviceServer *ServiceServer
var once sync.Once

func GetInstance() *ServiceServer {

	once.Do(
		func() {
			serviceServer = &ServiceServer{}
		},
	)

	return serviceServer
}

func InitGrpcSrv(wg *sync.WaitGroup) {

	config := config.GetInstance()

	sugar := logger.GetInstance()
	addr := config.GrpcServerAddress

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		sugar.Fatalf("Failed to listen on: %v", err)
	}

	sugar.Infof("Listening on %s", addr)

	serviceServer := GetInstance()
	grpcServer := grpc.NewServer()

	pb.RegisterGlobalEntryNotifyServiceServer(grpcServer, serviceServer)

	if err = grpcServer.Serve(lis); err != nil {
		sugar.Fatalf("Failed to serve: %v", err)
	}

	wg.Done()
}
