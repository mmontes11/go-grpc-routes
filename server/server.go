package server

import (
	ctx "context"
	"encoding/json"
	"errors"
	"log"
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
	for _, feature := range s.savedFeatures {
		if inRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *routeServer) RecordRoute(stream pb.Route_RecordRouteServer) error {
	return nil
}

func (s *routeServer) RouteChat(stream pb.Route_RouteChatServer) error {
	return nil
}

func (s *routeServer) loadFeatures() {
	if err := json.Unmarshal(data, &s.savedFeatures); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func newRouteServer() *routeServer {
	return &routeServer{}
}

// NewServer creates a new gRPC server
func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	routeServer := newRouteServer()
	routeServer.loadFeatures()
	pb.RegisterRouteServer(grpcServer, routeServer)
	return grpcServer
}
