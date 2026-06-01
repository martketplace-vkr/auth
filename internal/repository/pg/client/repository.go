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
			token_version,
			status_reason,
			status_updated_at,
			status_updated_by,
			created_at,
			updated_at
		from auth."user"
		where email = $1
	`

	user = new(domain.User)

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

func (r *repository) SelectUserByID(ctx context.Context, id int64) (user *domain.User, err error) {
	query := `
		select
			id,
			email,
			password_hash,
			email_verified,
			status,
			token_version,
			status_reason,
			status_updated_at,
			status_updated_by,
			created_at,
			updated_at
		from auth."user"
		where id = $1
	`

	user = new(domain.User)

	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		user,
		query,
		id,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) RecordActivity(ctx context.Context, userID int64) error {
	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(ctx, `
		insert into auth.client_activity_daily (client_id, activity_day, first_seen_at, last_seen_at)
		values ($1, (now() at time zone 'Europe/Moscow')::date, now(), now())
		on conflict (client_id, activity_day) do update
		set last_seen_at = excluded.last_seen_at
	`, userID)
	return err
}
