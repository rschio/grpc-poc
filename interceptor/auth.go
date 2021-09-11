package interceptor

import (
	"context"

	"github.com/rschio/grpc-poc/auth"
	"github.com/rschio/grpc-poc/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func getToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "failed to read context metadata")
	}

	tokens, ok := md["authorization"]
	if !ok || len(tokens) == 0 {
		return "", nil
	}
	return tokens[0], nil
}

func UnaryAuth(a auth.Auther) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		t, err := getToken(ctx)
		if err != nil {
			return nil, err
		}
		u, err := a.Auth(t, info.FullMethod)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		ctxUser := user.NewContext(ctx, u)
		return handler(ctxUser, req)
	}
}

type ssAuth struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *ssAuth) Context() context.Context {
	return ss.ctx
}

func StreamAuth(a auth.Auther) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := ss.Context()
		t, err := getToken(ctx)
		if err != nil {
			return err
		}
		u, err := a.Auth(t, info.FullMethod)
		if err != nil {
			return status.Error(codes.Unauthenticated, err.Error())
		}
		ctxUser := user.NewContext(ctx, u)
		newSS := &ssAuth{ServerStream: ss, ctx: ctxUser}
		return handler(srv, newSS)
	}
}
