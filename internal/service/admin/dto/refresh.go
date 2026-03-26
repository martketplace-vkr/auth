package dto

import adminpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/admin"

type (
	RefreshResponse struct {
		AccessToken  string
		RefreshToken string
	}
)

func (r *RefreshResponse) ToProto() *adminpb.RefreshTokenResponse {
	return &adminpb.RefreshTokenResponse{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}
