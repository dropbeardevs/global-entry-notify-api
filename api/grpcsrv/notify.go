package grpcsrv

import (
	"context"

	pb "bitbucket.org/dropbeardevs/global-entry-notify-api/api/proto"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
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

func (s *ServiceServer) GrpcUpdateNotificationToken(ctx context.Context, in *pb.UpdateNotificationTokenRq) (*pb.Error, error) {

	sugar := logger.GetInstance()
	sugar.Debugln("GrpcUpdateNotificationToken called")

	var response pb.Error

	err := notify.UpdateNotificationToken(in.UserId, in.Token)
	if err != nil {
		sugar.Error(err)
		response.Error = err.Error()

		return &response, err
	} else {
		response.Error = ""
	}

	return &response, nil
}

func (s *ServiceServer) GrpcUpdateNotificationDetails(ctx context.Context, in *pb.UpdateNotificationDetailsRq) (*pb.Error, error) {

	sugar := logger.GetInstance()
	sugar.Debugln("GrpcUpdateNotificationDetails called")

	var response pb.Error

	n := models.NotificationDetails{
		LocationId:      int(in.NotificationDetails.LocationId),
		TargetDate:      in.NotificationDetails.TargetDate.AsTime(),
		AppointmentDate: in.NotificationDetails.AppointmentDate.AsTime(),
	}

	err := notify.UpdateNotificationDetails(in.UserId, n)
	if err != nil {
		sugar.Error(err)
		response.Error = err.Error()

		return &response, err
	} else {
		response.Error = ""
	}

	return &response, nil
}

func (s *ServiceServer) GrpcDeleteNotificationDetails(ctx context.Context, in *pb.DeleteNotificationDetailsRq) (*pb.Error, error) {
	sugar := logger.GetInstance()
	sugar.Debugln("GrpcDeleteNotificationDetails called")

	var response pb.Error

	err := notify.DeleteNotificationDetails(in.UserId, int(in.LocationId))
	if err != nil {
		sugar.Error(err)
		response.Error = err.Error()

		return &response, err
	} else {
		response.Error = ""
	}

	return &response, nil
}
