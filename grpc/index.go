package grpc

import (
	"fmt"
	"net"

	gm "github.com/log-agent/emqx-grpc/grpc/proto"
	"google.golang.org/grpc"
)

type ManagerServer struct {
	gm.UnimplementedGrpcManagerServer
}

func GrpcRun() error {
	list, err := net.Listen("tcp", "0.0.0.0:18088")
	if err != nil {
		return fmt.Errorf("grpc linsten tcp error :%w", err)
	}
	var recvSize = 500 * 1024 * 1024
	var sendSize = 500 * 1024 * 1024
	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(recvSize),
		grpc.MaxSendMsgSize(sendSize),
	}
	s := grpc.NewServer(options...)
	gm.RegisterGrpcManagerServer(s, &ManagerServer{})
	if err := s.Serve(list); err != nil {
		return fmt.Errorf("grpc start error:%w", err)
	}
	return nil
}
