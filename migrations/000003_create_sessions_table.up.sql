CREATE TABLE sessions (
	id serial not null primary key,
	token varchar(50) not null,
	user_id integer not null,
	user_agent varchar(255) not null,
	created_at timestamp not null,
	expired_at timestamp not null
);
