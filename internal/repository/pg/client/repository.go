package client

import (
	"context"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
	"github.com/martketplace-vkr/auth/domain"
)

type repository struct {
	ctxGetter *trmsqlx.CtxGetter
	db        *sqlx.DB
}

func New(db *sqlx.DB, ctxGetter *trmsqlx.CtxGetter) *repository {
	return &repository{
		db:        db,
		ctxGetter: ctxGetter,
	}
}

func (r *repository) InsertUser(ctx context.Context, user *domain.User) error {
	query := `
		insert into auth."user"(
			email,
			password_hash
		) values (
			$1,
			$2 
		) returning *
	`

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		user,
		query,
		user.Email,
		user.PasswordHash,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SelectUserByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	query := `
		select 
			id,
			email,
			password_hash,
			email_verified,
			status,
			created_at,
			updated_at
		from auth."user"
		where email = $1
	`

	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		user,
		query,
		email,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
