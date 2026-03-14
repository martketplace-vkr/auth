package client

import (
	"context"

	"github.com/martketplace-vkr/auth/pkg/api/grpc/v1/client"
)

type Handler struct {
	service service
	client.UserClientServiceServer
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetUser(ctx context.Context, request *client.GetUserRequest) (resp *client.User, err error) {
	return resp, nil
}

func (h *Handler) UpdateUser(ctx context.Context, request *client.UpdateUserRequest) (resp *client.User, err error) {
	return resp, nil
}

func (h *Handler) CreateVendor(ctx context.Context, request *client.CreateVendorRequest) (resp *client.Vendor, err error) {
	return resp, nil
}

func (h *Handler) GetVendor(ctx context.Context, request *client.GetVendorRequest) (resp *client.Vendor, err error) {
	return resp, nil

}
