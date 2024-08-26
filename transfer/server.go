package transfer

import (
	"context"
	"errors"
	"github.com/lrayt/moving-bricks/dto/pb"
	"github.com/lrayt/moving-bricks/pkg/auth"
	"github.com/lrayt/moving-bricks/transfer/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type WrappedServerStream struct {
	grpc.ServerStream
	WrappedContext context.Context
}

func (w *WrappedServerStream) Context() context.Context {
	return w.WrappedContext
}

func WrapServerStream(stream grpc.ServerStream, user *auth.UserInfo) *WrappedServerStream {
	if existing, ok := stream.(*WrappedServerStream); ok {
		return existing
	}
	ctx := context.WithValue(stream.Context(), auth.UserId, user.UID)
	ctx = context.WithValue(ctx, auth.UserName, user.Name)
	return &WrappedServerStream{ServerStream: stream, WrappedContext: ctx}
}

type TServer struct {
	addr string
	svr  *grpc.Server
}

func (s TServer) auth(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok || len(md.Get(auth.Authorization)) <= 0 {
		return errors.New("miss authorization! ")
	}
	if user, err1 := auth.ParseToken(md.Get(auth.Authorization)[0]); err1 != nil {
		return err1
	} else {
		return handler(srv, WrapServerStream(ss, user))
	}
}

func (s *TServer) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.svr = grpc.NewServer(grpc.StreamInterceptor(s.auth))
	pb.RegisterMovingBricksServer(s.svr, new(handler.RPCHandler))
	return s.svr.Serve(listener)
}

func (s TServer) Stop() {
	s.svr.GracefulStop()
}

func NewServer(addr string) *TServer {
	return &TServer{addr: addr}
}
