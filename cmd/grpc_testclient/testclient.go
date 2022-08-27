package main

import (
	"context"
	"log"
	"os"
	"path"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "bitbucket.org/dropbeardevs/global-entry-notify-api/api/proto"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/initapp"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
)

var addr string = "localhost:5051"

func init() {
	// Get current running filename
	_, filename, _, _ := runtime.Caller(0)

	// Change to ../..
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	initapp.InitApp()

	sugar := logger.GetInstance()

	sugar.Infof("Running folder: %v", dir)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGlobalEntryNotifyServiceClient(conn)

	getNotificationDetails(c)
}

func getNotificationDetails(c pb.GlobalEntryNotifyServiceClient) {

	sugar := logger.GetInstance()
	sugar.Infoln("getNotificationDetails was invoked")

	getNotificationDetailsRq := pb.GetNotificationDetailsRq{
		UserId: "3D05A979-35F9-4A40-B075-444DEB63537A",
	}

	res, err := c.GrpcGetNotificationDetails(context.Background(), &getNotificationDetailsRq)

	if err != nil {
		sugar.Fatalf("Could not get notification details: %v", err)
	}

	for _, n := range res.NotificationDetails {
		sugar.Infof("Notification Details: %#v", n)
	}
}
