package server

import (
	ctx "context"
	"errors"
	"sync"

	"github.com/golang/protobuf/proto"
	pb "github.com/mmontes11/go-grpc-routes/route"
	"google.golang.org/grpc"
)

type routeServer struct {
	pb.UnimplementedRouteServer
	savedFeatures []*pb.Feature

	mux        sync.Mutex
	routeNotes map[string][]*pb.RouteNote
}

var errNotFound = errors.New("Not found")

func (s *routeServer) GetFeature(ctx ctx.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, errNotFound
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
