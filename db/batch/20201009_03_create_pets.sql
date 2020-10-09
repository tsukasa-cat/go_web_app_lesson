create table pets (
    id INT(64) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    category_id INT(64) NOT NULL,
    name VARCHAR(255) NOT NULL,
    status INT(32) NOT NULL,
    CONSTRAINT pets_fk_category_id
        FOREIGN KEY (category_id) 
        REFERENCES categories (id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB CHARSET=utf8;