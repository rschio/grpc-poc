package server

import (
	"context"
	"time"

	pb "github.com/rschio/grpc-poc/tracker/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedTrackerServer
}

func (s *Server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	const code = "BR000000000BR"
	if req.TrackingCode != code {
		return nil, status.Error(codes.NotFound, "parcel not found")
	}
	return &pb.SearchResponse{
		TrackingCode: code,
		Events: []*pb.Event{
			&pb.Event{Date: timestamppb.Now(), Status: "Vai chegar", Place: "Curitiba"},
			&pb.Event{
				Date:   timestamppb.New(time.Now().Add(2 * 24 * time.Hour)),
				Status: "Confia",
				Place:  "São Paulo",
			},
		},
	}, nil
}

func (s *Server) Login(ctx context.Context, creds *pb.Credentials) (*pb.Token, error) {
	if creds.User != "moises" || creds.Pass != "naoconsegue" {
		return nil, status.Error(codes.Unauthenticated, "wrong user or pass")
	}
	return &pb.Token{Value: "Bearer consegue", ExpiresIn: 10000000}, nil
}

func (s *Server) SearchBatch(req *pb.SearchBatchRequest, stream pb.Tracker_SearchBatchServer) error {
	resps := make([]*pb.SearchResponse, len(req.TrackingCodes))
	for i, code := range req.TrackingCodes {
		resps[i] = &pb.SearchResponse{TrackingCode: code}
	}

	for _, resp := range resps {
		if err := stream.Send(resp); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
