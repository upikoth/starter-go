CREATE TABLE registrations (
	id serial not null primary key,
	name varchar(20) not null,
	email varchar(255) not null unique,
	password_hash varchar(100) not null,
	registration_confirmation_token varchar(50) not null
);
