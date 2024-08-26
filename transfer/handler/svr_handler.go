package handler

import (
	"context"
	"fmt"
	"github.com/lrayt/moving-bricks/dto/pb"
	"github.com/lrayt/moving-bricks/pkg/auth"
	"github.com/lrayt/moving-bricks/pkg/string_util"
	"github.com/lrayt/moving-bricks/pkg/ts_error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"sync/atomic"
	"time"
)

type RPCHandler struct {
	UserNum uint32
	pb.UnimplementedMovingBricksServer
}

func (h *RPCHandler) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	atomic.AddUint32(&h.UserNum, 1)
	if token, err := auth.GenToken(&auth.UserInfo{UID: string_util.UUID(), Name: fmt.Sprintf("用户_%d", h.UserNum)}); err != nil {
		return nil, err
	} else {
		return &pb.PingResponse{Code: ts_error.OK, Msg: ts_error.Success, Delay: time.Now().Unix() - req.Timestamp, Token: token}, nil
	}
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
		//md, ok := metadata.FromIncomingContext(stream.Context())
		//if ok {
		log.Println(stream.Context().Value(auth.UserName), "svc:", chunk)
		//}
		chunk.Statue = pb.TransferStatus_TS_Success
		if err2 := stream.Send(chunk); err != nil {
			log.Println("err2:", err2)
		}
	}
	return nil
}
