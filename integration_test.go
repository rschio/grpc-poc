//go:build integration

package server

import (
	"context"
	"io"
	"log"
	"net"
	"testing"
	"time"

	"github.com/rschio/grpc-poc/server"
	pb "github.com/rschio/grpc-poc/tracker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterTrackerServer(s, new(server.Server))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestSearchBatch(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "fake",
		grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}

	c := pb.NewTrackerClient(conn)
	authCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer consegue")

	req := &pb.SearchBatchRequest{
		TrackingCodes: []string{"BR999999999BR", "BR888888888BR"},
	}

	t.Run("normal", func(t *testing.T) {
		stream, err := c.SearchBatch(authCtx, req)
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; ; i++ {
			rsp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatal(err)
			}
			got := rsp.TrackingCode
			want := req.TrackingCodes[i]
			if got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("timeout", func(t *testing.T) {
		ctx, _ := context.WithTimeout(authCtx, -1*time.Second)
		_, err := c.SearchBatch(ctx, req)

		if status.Code(err) != codes.DeadlineExceeded {
			t.Fatalf("got: %v, want dealineExceeded error", err)
		}

	})
}
