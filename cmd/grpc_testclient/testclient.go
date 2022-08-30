package main

import (
	"context"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "bitbucket.org/dropbeardevs/global-entry-notify-api/api/proto"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/initapp"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
)

var addr string = "localhost:8080"

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
	//updateNotificationToken(c)
	//getLocations(c)
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
		d := protojson.Format(n)
		sugar.Infof("Notification Details: %#v", d)
	}
}

func updateNotificationToken(c pb.GlobalEntryNotifyServiceClient) {

	sugar := logger.GetInstance()
	sugar.Infoln("updateNotificationToken was invoked")

	updateNotificationTokenRq := pb.UpdateNotificationTokenRq{
		UserId: "3D05A979-35F9-4A40-B075-444DEB63537A",
		Token:  uuid.NewString(),
	}

	_, err := c.GrpcUpdateNotificationToken(context.Background(), &updateNotificationTokenRq)

	if err != nil {
		sugar.Fatalf("Could not update notification token details: %v", err)
	}
}

func getLocations(c pb.GlobalEntryNotifyServiceClient) {

	sugar := logger.GetInstance()
	sugar.Infoln("getLocations was invoked")

	var empty emptypb.Empty

	pblocationList, err := c.GrpcGetLocations(context.Background(), &empty)
	if err != nil {
		sugar.Fatalf("Could not get locations, details: %v", err)
	}

	sugar.Infof("%#v", protojson.Format(pblocationList))
}
