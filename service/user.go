package service

import (
	"strconv"

	"github.com/bananafried525/wallet/user/database"
	"github.com/bananafried525/wallet/user/entity"
)

type GetUserRequest struct {
	ID string `validate:"required,number"`
}

type GetUserResponse struct {
	ID           uint32
	UserName     string
	Email        string
	ReferralCode string
}

func GetUser(db *database.Connection, req *GetUserRequest) *GetUserResponse {
	dbTxn := db.Begin()

	userId, err := strconv.ParseUint(req.ID, 10, 32)
	if err != nil {
		dbTxn.Rollback()
		return nil
	}

	user, err := entity.GetUser(dbTxn, uint32(userId))
	if err != nil {
		dbTxn.Rollback()
		return nil
	}

	dbTxn.Commit()

	return &GetUserResponse{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		ReferralCode: user.ReferralCode,
	}
}
