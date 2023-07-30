CREATE TABLE users (
	id serial not null primary key,
	name varchar(20) not null,
	email varchar(255) not null unique,
	status varchar(10) not null,
	password_hash varchar(100) not null
);
