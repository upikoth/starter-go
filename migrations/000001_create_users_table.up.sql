CREATE TABLE users (
	id serial not null primary key,
	email varchar(255) not null unique,
	password_hash varchar(100) not null
);
