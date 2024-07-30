package user_pb

import (
	"context"

	"github.com/bananafried525/wallet/user/database"
	"github.com/bananafried525/wallet/user/service"
)

type userServer struct {
	UnimplementedUserServiceServer
	db *database.Connection
}

func New(db *database.Connection) userServer {
	return userServer{
		db: db,
	}
}

func (us userServer) GetUser(ctx context.Context, u *GetUserRequest) (*GetUserResponse, error) {

	result := service.GetUser(us.db, service.GetUserRequest{ID: u.GetId()})

	return &GetUserResponse{
		User: &GetUserResponse_User{
			Id:           result.ID,
			UserName:     result.UserName,
			Email:        result.Email,
			ReferralCode: result.ReferralCode,
		}}, nil
}

func (us userServer) CreateUser(ctx context.Context, u *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, nil
}
