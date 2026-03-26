package admin

import (
	"context"
	"database/sql"
	"errors"

	"github.com/martketplace-vkr/auth/internal/service/admin/dto"
)

func (s *service) SignUp(ctx context.Context, req dto.SignUpRequest) (resp dto.SignUpResponse, err error) {
	hasInvite, err := s.repository.HasActiveInviteToken(ctx, req.InviteToken)
	if err != nil {
		return resp, err
	}
	if !hasInvite {
		return resp, ErrInviteTokenInvalid
	}

	req.Password, err = hashPassword(req.Password)
	if err != nil {
		return resp, err
	}

	user := req.ToDomain()

	err = s.txManager.Do(ctx, func(ctx context.Context) error {
		if err := s.repository.InsertUser(ctx, user); err != nil {
			return err
		}

		if err := s.repository.UseInviteToken(ctx, req.InviteToken, user.ID); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ErrInviteTokenInvalid
			}

			return err
		}

		return nil
	})
	if err != nil {
		return resp, err
	}

	resp.UserID = user.ID

	return resp, nil
}
