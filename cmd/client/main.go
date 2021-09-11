package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/rschio/grpc-poc/tracker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	const port = "9090"
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewTrackerClient(conn)

	token, err := c.Login(ctx, &pb.Credentials{User: "moises", Pass: "naoconsegue"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token: %v\n", token)

	authCtx := metadata.AppendToOutgoingContext(ctx, "authorization", token.Value)

	rsp, err := c.Search(authCtx, &pb.SearchRequest{TrackingCode: "BR000000000BR"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Search Response: %v\n", rsp)

	stream, err := c.SearchBatch(authCtx, &pb.SearchBatchRequest{
		TrackingCodes: []string{
			"BR000000000BR",
			"BR111111111BR",
			"BR222222222BR",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		rsp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Search Response: %v\n", rsp)
	}
}
