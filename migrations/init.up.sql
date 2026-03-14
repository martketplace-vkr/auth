
create schema if not exists auth;

create table if not exists auth."user" (
    id bigserial primary key,
    email text not null unique,
    password_hash text not null,
    email_verified boolean not null default false,
    status text not null default 'active',
    created_at timestamptz not null default now(),
    updated_at timestamptz
)

create table if not exists auth.sessions (
    id bigserial primary key,
    user_id bigint not null references auth.users(id),
    ip_address inet,
    user_agent text,
    created_at timestamptz not null default now(),
    last_seen timestamptz
);