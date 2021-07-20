CREATE TABLE IF NOT EXISTS contactlist(
    id primary key serial,
    name varchar(80) not null,
	phone varchar(13) not null,
	gender varchar(6) not null,
	email varchar(50) not null,
	createat timestamp default now(),
)
create table tasklist(
	id primary key serial,
	assignee not null text,
	title not null text,
	deadline not null timestamp,
	done boolean,
)