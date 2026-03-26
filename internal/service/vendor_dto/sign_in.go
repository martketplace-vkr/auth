package vendordto

import vendorpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/vendor"

type (
	SignInRequest struct {
		Email    string
		Password string
	}
	SignInResponse struct {
		AccessToken  string
		RefreshToken string
	}
)

func SignInRequestFromProto(req *vendorpb.LoginRequest) SignInRequest {
	return SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (r *SignInResponse) ToProto() *vendorpb.LoginResponse {
	return &vendorpb.LoginResponse{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}
