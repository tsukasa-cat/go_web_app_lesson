create table users (
    id INT(64) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255),
    auth_token TEXT
) ENGINE=InnoDB CHARSET=utf8;