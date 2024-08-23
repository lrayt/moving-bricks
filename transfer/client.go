package transfer

import (
	"context"
	"github.com/lrayt/moving-bricks/dto/pb"
	"github.com/lrayt/moving-bricks/pkg/string_util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

type TClient struct {
	rpc pb.MovingBricksClient
	ctx context.Context
}

func NewClient(addr string) (*TClient, error) {
	// dial
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	// client
	rpc := pb.NewMovingBricksClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()
	res, err1 := rpc.Ping(ctx, &pb.PingRequest{Timestamp: time.Now().Unix()})
	if err1 != nil {
		return nil, err1
	}
	md := metadata.Pairs("authorization", "Bearer "+res.Token)
	ctx1 := metadata.NewOutgoingContext(context.Background(), md)
	return &TClient{rpc: rpc, ctx: ctx1}, nil
}

func (c TClient) Transfer() error {
	stream, err := c.rpc.Transfer(c.ctx)
	if err != nil {
		return err
	}
	go func() {
		for {
			chunk, err1 := stream.Recv()
			if err1 != nil {
				log.Fatalf("recv err:%s", err1.Error())
			} else {
				log.Println(chunk)
			}
		}
	}()
	for i := 0; i < 5; i++ {
		if err2 := stream.Send(&pb.TransferChunk{TcId: string_util.UUID()}); err2 != nil {
			log.Fatalf("err2:%s", err2.Error())
		}
		time.Sleep(time.Second * 3)
	}
	return stream.CloseSend()
}
