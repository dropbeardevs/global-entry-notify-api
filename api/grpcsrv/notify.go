package grpcsrv

import (
	"context"

	pb "bitbucket.org/dropbeardevs/global-entry-notify-api/api/proto"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/notify"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ServiceServer) GrpcGetNotificationDetails(ctx context.Context, in *pb.GetNotificationDetailsRq) (*pb.GetNotificationDetailsRp, error) {

	sugar := logger.GetInstance()
	sugar.Debugln("GetNotificationDetails called")

	var response pb.GetNotificationDetailsRp

	notificationDetails, err := notify.GetNotificationDetails(in.UserId)
	if err != nil {
		sugar.Error(err)
		return nil, err
	}

	for _, notificationDetail := range notificationDetails {

		var n pb.NotificationDetails

		n.LocationId = int32(notificationDetail.LocationId)
		n.TargetDate = timestamppb.New(notificationDetail.TargetDate)
		n.AppointmentDate = timestamppb.New(notificationDetail.AppointmentDate)
		n.LastNotifiedDate = timestamppb.New(notificationDetail.LastNotifiedDate)

		response.NotificationDetails = append(response.NotificationDetails, &n)
	}

	return &response, nil

}
