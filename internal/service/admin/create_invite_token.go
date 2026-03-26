package admin

import (
	"context"

	"github.com/martketplace-vkr/auth/internal/service/admin/dto"
)

func (s *service) CreateInviteToken(
	ctx context.Context,
	req dto.CreateInviteTokenRequest,
) (resp dto.CreateInviteTokenResponse, err error) {
	token, err := generateInviteToken()
	if err != nil {
		return resp, err
	}

	err = s.repository.CreateInviteToken(ctx, req.UserID, req.RoleID, token)
	if err != nil {
		return resp, err
	}

	resp.Token = token

	return resp, nil
}
