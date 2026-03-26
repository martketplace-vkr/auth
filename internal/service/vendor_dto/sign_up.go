package vendordto

import (
	"github.com/martketplace-vkr/auth/domain"
	vendorpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/vendor"
)

type (
	SignUpRequest struct {
		Email    string
		Password string
	}
	SignUpResponse struct {
		VendorID int64
	}
)

func SignUpRequestFromProto(req *vendorpb.RegisterRequest) SignUpRequest {
	return SignUpRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (r *SignUpRequest) ToDomain() *domain.User {
	return &domain.User{
		Email:        r.Email,
		PasswordHash: r.Password,
	}
}

func (r *SignUpResponse) ToProto() *vendorpb.RegisterResponse {
	return &vendorpb.RegisterResponse{
		VendorId: r.VendorID,
	}
}
