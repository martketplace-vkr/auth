package client

import (
	"context"

	"github.com/martketplace-vkr/auth/internal/service/client/dto"
)

func (s *service) SignUp(ctx context.Context, req dto.SignUpRequest) (resp dto.SignUpResponse, err error) {
	req.Password, err = hashPassword(req.Password)
	if err != nil {
		return resp, err
	}

	user := req.ToDomain()

	err = s.repository.InsertUser(ctx, user)
	if err != nil {
		return resp, err
	}

	resp.UserId = user.ID

	return resp, nil
}
