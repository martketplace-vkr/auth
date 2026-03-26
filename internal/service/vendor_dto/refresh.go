package vendordto

import vendorpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/vendor"

type (
	RefreshResponse struct {
		AccessToken  string
		RefreshToken string
	}
)

func (r *RefreshResponse) ToProto() *vendorpb.RefreshTokenResponse {
	return &vendorpb.RefreshTokenResponse{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}
