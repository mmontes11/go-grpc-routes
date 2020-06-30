package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/mmontes11/go-grpc-routes/route"
	"google.golang.org/grpc"
)

var (
	server         = flag.String("server", "localhost:10000", "The server address in the format of host:port")
	timeoutSeconds = flag.Int("timeout", 10, "Request timeout in seconds")
	timeout        = time.Duration(*timeoutSeconds) * time.Second
)

func getFeature(client pb.RouteClient, point *pb.Point) {
	log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	feature, err := client.GetFeature(ctx, point)
	if err != nil {
		log.Printf("Error getting feature: %v", err)
		return
	}
	log.Println(feature)
}

func listFeatures(client pb.RouteClient, rect *pb.Rectangle) {
	log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Printf("Error listing features: %v", err)
		return
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Printf("Error listing features: %v", err)
			break
		}
		log.Println(feature)
	}
}

func main() {
	flag.Parse()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(timeout),
	}
	conn, err := grpc.Dial(*server, opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteClient(conn)

	getFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
	getFeature(client, &pb.Point{Latitude: 0, Longitude: 0})

	listFeatures(client, &pb.Rectangle{
		Lo: &pb.Point{Latitude: 410000000, Longitude: -740000000},
		Hi: &pb.Point{Latitude: 415000000, Longitude: -745000000},
	})
}
