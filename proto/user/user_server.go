package user_pb

import (
	"context"

	"github.com/bananafried525/wallet/user/database"
	"github.com/bananafried525/wallet/user/service"
	"github.com/bananafried525/wallet/user/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	data, err := validator.GetUser(map[string]string{"id": u.GetId()})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	result := service.GetUser(us.db, data)

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
