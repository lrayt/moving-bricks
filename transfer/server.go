package transfer

import (
	"github.com/lrayt/moving-bricks/dto/pb"
	"github.com/lrayt/moving-bricks/transfer/handler"
	"google.golang.org/grpc"
	"net"
)

type TServer struct {
	listener net.Listener
	svr      *grpc.Server
}

func (s TServer) Run() error {
	pb.RegisterMovingBricksServer(s.svr, new(handler.RPCHandler))
	return s.svr.Serve(s.listener)
}

func (s TServer) Stop() {
	s.svr.GracefulStop()
}

func NewServer(addr string) (*TServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &TServer{listener: listener, svr: grpc.NewServer(
		//grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		//	md, ok := metadata.FromIncomingContext(ctx)
		//	if !ok {
		//		log.Println("not ok---->")
		//	} else {
		//		log.Println(md.Get("authorization"))
		//	}
		//	ctx = context.WithValue(ctx, "uid", "admin")
		//	// 继续处理请求
		//	return handler(ctx, req)
		//}),
		grpc.StreamServerInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return handler(srv, ss)
		}),
		//grpc.StreamInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		//	md, ok := metadata.FromIncomingContext(ss.Context())
		//	if !ok {
		//		log.Println("not ok---->")
		//	} else {
		//		log.Println("-->", md.Get("authorization"))
		//	}
		//	//metadata.AppendToOutgoingContext(ss.Context(), "uid", "admin")
		//	ss.SetHeader(metadata.New(map[string]string{"uid": "admin"}))
		//	return handler(srv, ss)
		//}),
	)}, nil
}
