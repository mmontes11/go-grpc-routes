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
			log.Printf("Error receiving features: %v", err)
			break
		}
		log.Println(feature)
	}
}

func recordRoute(client pb.RouteClient) {
	points := randomPoints()
	log.Printf("Transversing %d points", len(points))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Printf("Error recording route: %v", err)
		return
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Printf("Error sending points: %v", err)
			return
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error receiving reply: %v", err)
		return
	}
	log.Printf("Route summary: %v", reply)
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

	getFeature(client, validPoint)
	getFeature(client, invalidPoint)
	listFeatures(client, rect)
	recordRoute(client)
}
