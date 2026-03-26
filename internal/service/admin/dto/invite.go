package dto

import adminpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/admin"

type (
	CreateInviteTokenRequest struct {
		UserID int64
		RoleID int64
	}
	CreateInviteTokenResponse struct {
		Token string
	}
)

func CreateInviteTokenRequestFromProto(req *adminpb.CreateInviteTokenRequest) CreateInviteTokenRequest {
	return CreateInviteTokenRequest{
		UserID: req.UserId,
		RoleID: req.RoleId,
	}
}

func (r *CreateInviteTokenResponse) ToProto() *adminpb.CreateInviteTokenResponse {
	return &adminpb.CreateInviteTokenResponse{
		Token: r.Token,
	}
}
