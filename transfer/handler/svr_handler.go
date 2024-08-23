package handler

import (
	"context"
	"github.com/lrayt/moving-bricks/dto/pb"
	"github.com/lrayt/moving-bricks/pkg/string_util"
	"github.com/lrayt/moving-bricks/pkg/ts_error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

type RPCHandler struct {
	pb.UnimplementedMovingBricksServer
}

func (R RPCHandler) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Code:  ts_error.OK,
		Msg:   ts_error.Success,
		Delay: time.Now().Unix() - req.Timestamp,
		Token: string_util.UUID(),
	}, nil
}

func (R RPCHandler) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (R RPCHandler) Transfer(stream pb.MovingBricks_TransferServer) error {
	for {
		chunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("complate!")
			} else if status.Code(err) == codes.Canceled {
				log.Println("client cancel or close")
			} else {
				log.Println("---->err:", err)
			}
			break
		}
		md, ok := metadata.FromIncomingContext(stream.Context())
		if ok {
			log.Println(md.Get("uid"), "svc:", chunk)
		}
		chunk.Statue = pb.TransferStatus_TS_Success
		if err2 := stream.Send(chunk); err != nil {
			log.Println("err2:", err2)
		}
	}
	return nil
}
