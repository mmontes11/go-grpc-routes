package server

import (
	ctx "context"

	pb "github.com/mmontes11/go-grpc-routes/route"
	"google.golang.org/grpc"
)

type routeServer struct {
}

func (s *routeServer) GetFeature(ctx ctx.Context, point *pb.Point) (*pb.Feature, error) {
	return nil, nil
}

func (s *routeServer) ListFeatures(rect *pb.Rectangle, stream pb.Route_ListFeaturesServer) error {
	return nil
}

func (s *routeServer) RecordRoute(stream pb.Route_RecordRouteServer) error {
	return nil
}

func (s *routeServer) RouteChat(stream pb.Route_RouteChatServer) error {
	return nil
}

func newRouteServer() *routeServer {
	return &routeServer{}
}

// NewServer creates a new gRPC server
func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterRouteServer(grpcServer, newRouteServer())
	return grpcServer
}
