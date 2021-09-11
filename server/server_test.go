package server

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	pb "github.com/rschio/grpc-poc/tracker/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLogin(t *testing.T) {
	s := new(Server)
	ctx := context.Background()
	t.Run("correct creds", func(t *testing.T) {
		creds := &pb.Credentials{User: "moises", Pass: "naoconsegue"}
		want := &pb.Token{Value: "Bearer consegue", ExpiresIn: 10000000}

		got, err := s.Login(ctx, creds)
		if err != nil {
			t.Error(err)
		}
		ignoreUnexported := cmpopts.IgnoreUnexported(pb.Token{})
		if !cmp.Equal(got, want, ignoreUnexported) {
			t.Error(cmp.Diff(got, want, ignoreUnexported))
		}
	})
	t.Run("incorrect creds", func(t *testing.T) {
		creds := &pb.Credentials{User: "moises", Pass: "consegue"}
		want := codes.Unauthenticated

		_, err := s.Login(ctx, creds)
		if err == nil {
			t.Fatal("got nil, want err")
		}
		if got := status.Convert(err).Code(); got != want {
			t.Fatalf("got code: %v, want: %v", got, want)
		}

	})
}
