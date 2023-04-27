-- Creating a table named ucc_users with columns uid, name, username, password, email, gender
create table ucc_users (
	uid INT PRIMARY KEY
	name VARCHAR(70) NOT NULL,
	username VARCHAR(20) NOT NULL UNIQUE,
	password VARCHAR(20) NOT NULL,
	email VARCHAR(50) NOT NULL,
	gender VARCHAR(1) NOT NULL,
	CHECK (gender in ('M', 'F'))
);

-- Inserting values into ucc_users table
insert into ucc_users (name, email, gender, username, password) values (0, 'Dummy Duck', 'dummy', 'dumdum', 'dummy@pace.edu', 'M');
insert into ucc_users (name, email, gender, username, password) values (1, 'Buzz Lightyear', 'buzz', 'busybee', 'buzz@pace.edu', 'M');
insert into ucc_users (name, email, gender, username, password) values (2, 'Snow White', 'snow', 'abc1234', 'snow@pace.edu', 'F');
