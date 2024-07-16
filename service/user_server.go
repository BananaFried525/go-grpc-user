package service

import (
	pb "bananaf/wallet/user/proto/user"
	"context"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() userServer {
	return userServer{}
}

func (userServer) GetUser(ctx context.Context, u *pb.User) (*pb.User, error) {
	return u, nil
}
