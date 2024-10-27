-- +goose Up
create table if not exists users (
     id string,
     email string,
     password_hash string,
     role string,
     primary key (id)
);

create table if not exists sessions (
    id string,
    token string,
    user_id string,
    primary key (id)
);

create table if not exists registrations (
     id string,
     email string,
     confirmation_token string,
     primary key (id)
);

create table if not exists password_recovery_requests (
     id string,
     email string,
     confirmation_token string,
     primary key (id)
);

-- +goose Down
drop table if exists users;
drop table if exists sessions;
drop table if exists registrations;
drop table if exists password_recovery_requests;
