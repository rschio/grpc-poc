package interceptor

import (
	"context"
	"log"
	"time"

	"github.com/rschio/grpc-poc/user"
	"google.golang.org/grpc"
)

func logUserAndTime(ctx context.Context, method string, start time.Time) {
	u, ok := user.FromContext(ctx)
	if ok {
		log.Printf("request from user: %s", u.Name)
	}
	log.Printf("duration of %s: %v", method, time.Since(start))
}

func UnaryLog(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	now := time.Now()
	defer logUserAndTime(ctx, info.FullMethod, now)
	return handler(ctx, req)
}

func StreamLog(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	now := time.Now()
	defer logUserAndTime(ss.Context(), info.FullMethod, now)
	return handler(srv, ss)
}
