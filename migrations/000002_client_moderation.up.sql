alter table auth."user"
    add column if not exists token_version bigint not null default 0,
    add column if not exists status_reason text,
    add column if not exists status_updated_at timestamptz,
    add column if not exists status_updated_by bigint;

create index if not exists auth_user_status_created_idx
    on auth."user" (status, created_at desc);

create index if not exists auth_user_email_lower_idx
    on auth."user" (lower(email));

create table if not exists auth.client_moderation_event (
    id bigserial primary key,
    client_id bigint not null references auth."user" (id),
    admin_id bigint not null,
    old_status text not null,
    new_status text not null,
    reason text not null,
    created_at timestamptz not null default now()
);

create index if not exists client_moderation_event_client_created_idx
    on auth.client_moderation_event (client_id, created_at desc);

create table if not exists auth.client_activity_daily (
    client_id bigint not null references auth."user" (id),
    activity_day date not null,
    first_seen_at timestamptz not null default now(),
    last_seen_at timestamptz not null default now(),
    primary key (client_id, activity_day)
);

create index if not exists client_activity_daily_day_idx
    on auth.client_activity_daily (activity_day);
