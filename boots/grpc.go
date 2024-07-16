package boots

import (
	"bananaf/wallet/user/config"
	userProto "bananaf/wallet/user/proto/user"
	"bananaf/wallet/user/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	cfg *config.Config
}

func NewGrpcServer(cfg *config.Config) *grpcServer {
	return &grpcServer{
		cfg: cfg,
	}
}

func (s *grpcServer) Start() error {
	otps := make([]grpc.ServerOption, 0)
	gs := grpc.NewServer(otps...)

	userProto.RegisterUserServiceServer(gs, service.NewUserServer())

	log.Printf("start grpc server on port %d", 50051)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs.Serve(listener)
	return nil
}
