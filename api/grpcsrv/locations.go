package grpcsrv

import (
	"context"
	"encoding/json"

	pb "bitbucket.org/dropbeardevs/global-entry-notify-api/api/proto"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/locations"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServiceServer) GrpcGetLocations(ctx context.Context, e *emptypb.Empty) (*pb.GetLocationsRp, error) {

	sugar := logger.GetInstance()
	sugar.Debugln("GrpcGetLocations called")

	response := &pb.GetLocationsRp{}

	locationList, err := locations.GetLocations()
	if err != nil {
		sugar.Error(err)
		return response, err
	}

	for _, loc := range *locationList {
		locJson, err := json.Marshal(loc)
		if err != nil {
			sugar.Error(err)
			return response, err
		}

		pbLoc := &pb.Location{}

		err = protojson.Unmarshal(locJson, pbLoc)
		if err != nil {
			sugar.Error(err)
			return response, err
		}

		response.LocationList = append(response.LocationList, pbLoc)
	}

	return response, nil
}
