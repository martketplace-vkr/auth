package vendor

import (
	"context"

	clientdto "github.com/martketplace-vkr/auth/internal/service/client/dto"
	vendordto "github.com/martketplace-vkr/auth/internal/service/vendor_dto"
)

func (s *service) SignIn(ctx context.Context, req vendordto.SignInRequest) (resp vendordto.SignInResponse, err error) {
	user, err := s.repository.SelectUserByEmail(ctx, req.Email)
	if err != nil {
		return resp, err
	}

	err = comparePassword(user.PasswordHash, req.Password)
	if err != nil {
		return resp, err
	}

	resp.AccessToken, err = s.generateAccessToken(user.ID)
	if err != nil {
		return resp, err
	}

	refreshToken, refreshHash, err := generateRefreshToken()
	if err != nil {
		return resp, err
	}

	resp.RefreshToken = refreshToken

	err = s.cache.SaveRefreshToken(ctx, clientdto.SaveRefreshTokenArgs{
		RefreshHash: refreshHash,
		UserID:      user.ID,
	})
	if err != nil {
		return resp, err
	}

	return resp, nil
}
