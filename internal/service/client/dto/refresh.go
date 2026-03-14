package dto

import "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/client"

type (
	RefreshResponse struct {
		AccessToken  string
		RefreshToken string
	}
)

func (r *RefreshResponse) ToProto() *client.RefreshTokenResponse {
	return &client.RefreshTokenResponse{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}
