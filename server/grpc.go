package server

import (
	"fmt"
	"log"
	"net"

	"github.com/bananafried525/wallet/user/config"
	"github.com/bananafried525/wallet/user/database"

	userProto "github.com/bananafried525/wallet/user/proto/user"

	"google.golang.org/grpc"
)

type grpcServer struct {
	cfg *config.Config
	db  *database.Connection
}

func New(cfg *config.Config, db *database.Connection) *grpcServer {
	return &grpcServer{
		cfg: cfg,
		db:  db,
	}
}

func (s *grpcServer) Start() error {
	otp := make([]grpc.ServerOption, 0)
	gs := grpc.NewServer(otp...)

	// register grpc service
	userProto.RegisterUserServiceServer(gs, userProto.New(s.db))

	log.Printf("start grpc server on port %d", s.cfg.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gs.Serve(listener)
	return nil
}
