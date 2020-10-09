create table users (
    id INT(64) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    phone VARCHAR(255),
    user_status INT(64),
    UNIQUE (username)
) ENGINE=InnoDB CHARSET=utf8;