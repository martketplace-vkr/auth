
create schema if not exists auth;

create table if not exists auth."user" (
    id bigserial primary key,
    email text not null unique,
    password_hash text not null,
    email_verified boolean not null default false,
    status text not null default 'active',
    created_at timestamptz not null default now(),
    updated_at timestamptz
);

create schema if not exists vendor;

create table if not exists vendor.vendor(
    id bigserial primary key,
    email text not null unique,
    password_hash text not null,
    email_verified boolean not null default false,
    status text not null default 'active',
    created_at timestamptz not null default now(),
    updated_at timestamptz
);

create schema if not exists employee;

create table if not exists employee.role(
    id bigserial primary key,
    role text not null unique
);

create table if not exists employee."user"(
    id bigserial primary key,
    email text not null unique,
    password_hash text not null,
    email_verified boolean not null default false,
    status text not null default 'active',
    created_at timestamptz not null default now(),
    updated_at timestamptz
);

create table if not exists employee.invite_token(
    id bigserial primary key,
    token text not null unique,
    created_by integer not null references employee."user"(id),
    used_by integer references employee."user"(id),
    role_id integer references employee.role(id),
    created_at timestamptz not null default now(),
    used_at timestamptz
);

create table if not exists employee.user_role(
    id bigserial primary key,
    employee_id integer not null references employee."user"(id),
    role_id integer not null references employee.role(id)
);