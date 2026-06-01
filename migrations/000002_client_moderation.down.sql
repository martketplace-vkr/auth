drop table if exists auth.client_activity_daily;
drop table if exists auth.client_moderation_event;

drop index if exists auth.auth_user_email_lower_idx;
drop index if exists auth.auth_user_status_created_idx;

alter table auth."user"
    drop column if exists status_updated_by,
    drop column if exists status_updated_at,
    drop column if exists status_reason,
    drop column if exists token_version;
