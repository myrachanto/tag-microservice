-- working with many to many relationships
-- user table
CREATE TABLE user (
	id bigint(20) NOT NULL, 
	email varchar(40) NOT NULL,
	username varchar(15) NOT NULL,
	password varchar(100) NOT NULL,
	PRIMARY KEY (id)
);
-- role table
CREATE TABLE role (
	id bigint(20) NOT NULL,
	name varchar(60) NOT NULL, 
	PRIMARY KEY (id)
);
-- user-roles for many to many roles ie one person can have many roles 
-- as well as one role can be assinged to many people
-- the norm of naming a junction able is use both table names seperated by an underscore
CREATE TABLE user_roles (
	user_id bigint(20) NOT NULL,
	role_id bigint(20) NOT NULL,
	FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE RESTRICT ON UPDATE CASCADE,
	FOREIGN KEY (role_id) REFERENCES role (id) ON DELETE RESTRICT ON UPDATE CASCADE,
	PRIMARY KEY (user_id, role_id)
);

-- quering data that has junction table 
SELECT user.id, user.email, user.username, role.id AS role_id, role.name AS role_name
FROM user 
JOIN user_roles on (user.id=user_roles.user_id)
JOIN role on (role.id=user_roles.role_id);