package domain

import "time"

type User struct {
	ID              int64      `db:"id"`
	Email           string     `db:"email"`
	PasswordHash    string     `db:"password_hash"`
	EmailVerified   bool       `db:"email_verified"`
	Status          string     `db:"status"`
	TokenVersion    int64      `db:"token_version"`
	StatusReason    *string    `db:"status_reason"`
	StatusUpdatedAt *time.Time `db:"status_updated_at"`
	StatusUpdatedBy *int64     `db:"status_updated_by"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
}

type ClientModerationEvent struct {
	ID        int64     `db:"id"`
	ClientID  int64     `db:"client_id"`
	AdminID   int64     `db:"admin_id"`
	OldStatus string    `db:"old_status"`
	NewStatus string    `db:"new_status"`
	Reason    string    `db:"reason"`
	CreatedAt time.Time `db:"created_at"`
}

type Client struct {
	User
	LastActivityAt   *time.Time `db:"last_activity_at"`
	ModerationEvents []ClientModerationEvent
}
