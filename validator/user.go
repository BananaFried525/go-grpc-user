package validator

import (
	"github.com/bananafried525/wallet/user/service"
	validate "github.com/go-playground/validator/v10"
)

func GetUser(params map[string]string) (*service.GetUserRequest, error) {
	data := &service.GetUserRequest{
		ID: params["id"],
	}

	v := validate.New(validate.WithRequiredStructEnabled())

	err := v.Struct(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
