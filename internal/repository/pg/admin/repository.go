package admin

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

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
		insert into employee."user"(
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
		from employee."user"
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
			created_at,
			updated_at
		from employee."user"
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

func (r *repository) HasActiveInviteToken(ctx context.Context, token string) (exists bool, err error) {
	query := `
		select exists(
			select 1
			from employee.invite_token
			where token = $1
				and used_at is null
		)
	`

	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		&exists,
		query,
		token,
	)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *repository) UseInviteToken(ctx context.Context, token string, userID int64) error {
	query := `
		update employee.invite_token
		set
			used_by = $2,
			used_at = $3
		where token = $1
			and used_at is null
	`

	result, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		query,
		token,
		userID,
		time.Now(),
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *repository) CreateInviteToken(
	ctx context.Context,
	createdBy int64,
	roleID int64,
	token string,
) error {
	query := `
		insert into employee.invite_token(
			token,
			created_by,
			role_id
		) values (
			$1,
			$2,
			$3
		)
	`

	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		query,
		token,
		createdBy,
		roleID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ListVendors(ctx context.Context) ([]domain.User, error) {
	query := `
		select
			id,
			email,
			password_hash,
			email_verified,
			status,
			created_at,
			updated_at
		from vendor.vendor
		order by email
	`

	var vendors []domain.User
	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).SelectContext(ctx, &vendors, query)
	if err != nil {
		return nil, err
	}

	return vendors, nil
}

func IsInviteNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func (r *repository) ListClients(ctx context.Context, search string, userStatus string, limit uint32, offset uint32) ([]domain.Client, uint64, error) {
	if limit == 0 || limit > 100 {
		limit = 50
	}
	search = strings.TrimSpace(search)
	userStatus = strings.TrimSpace(userStatus)

	var total uint64
	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(ctx, &total, `
		select count(*)
		from auth."user"
		where ($1 = '' or lower(email) like '%' || lower($1) || '%')
		  and ($2 = '' or status = $2)
	`, search, userStatus)
	if err != nil {
		return nil, 0, err
	}

	var clients []domain.Client
	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).SelectContext(ctx, &clients, `
		select
			u.id,
			u.email,
			u.password_hash,
			u.email_verified,
			u.status,
			u.token_version,
			u.status_reason,
			u.status_updated_at,
			u.status_updated_by,
			u.created_at,
			u.updated_at,
			a.last_activity_at
		from auth."user" u
		left join lateral (
			select max(last_seen_at) as last_activity_at
			from auth.client_activity_daily
			where client_id = u.id
		) a on true
		where ($1 = '' or lower(u.email) like '%' || lower($1) || '%')
		  and ($2 = '' or u.status = $2)
		order by u.created_at desc
		limit $3 offset $4
	`, search, userStatus, limit, offset)
	return clients, total, err
}

func (r *repository) GetClient(ctx context.Context, clientID int64) (domain.Client, error) {
	var client domain.Client
	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(ctx, &client, `
		select
			u.id,
			u.email,
			u.password_hash,
			u.email_verified,
			u.status,
			u.token_version,
			u.status_reason,
			u.status_updated_at,
			u.status_updated_by,
			u.created_at,
			u.updated_at,
			a.last_activity_at
		from auth."user" u
		left join lateral (
			select max(last_seen_at) as last_activity_at
			from auth.client_activity_daily
			where client_id = u.id
		) a on true
		where u.id = $1
	`, clientID)
	if err != nil {
		return client, err
	}
	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).SelectContext(ctx, &client.ModerationEvents, `
		select id, client_id, admin_id, old_status, new_status, reason, created_at
		from auth.client_moderation_event
		where client_id = $1
		order by created_at desc
	`, clientID)
	return client, err
}

func (r *repository) UpdateClientStatus(ctx context.Context, clientID int64, userStatus string, reason string, adminID int64) (domain.Client, error) {
	var oldStatus string
	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(ctx, &oldStatus, `
		select status from auth."user" where id = $1 for update
	`, clientID)
	if err != nil {
		return domain.Client{}, err
	}
	_, err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(ctx, `
		update auth."user"
		set status = $2,
		    status_reason = $3,
		    status_updated_at = now(),
		    status_updated_by = $4,
		    token_version = token_version + 1,
		    updated_at = now()
		where id = $1
	`, clientID, userStatus, reason, adminID)
	if err != nil {
		return domain.Client{}, err
	}
	_, err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(ctx, `
		insert into auth.client_moderation_event (client_id, admin_id, old_status, new_status, reason)
		values ($1, $2, $3, $4, $5)
	`, clientID, adminID, oldStatus, userStatus, reason)
	if err != nil {
		return domain.Client{}, err
	}
	return r.GetClient(ctx, clientID)
}
