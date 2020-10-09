create table orders (
    id INT(64) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    pet_id INT(64) NOT NULL,
    quantity INT(32),
    ship_date DATETIME,
    status INT(32) NOT NULL,
    complete BOOLEAN NOT NULL,
    CONSTRAINT orders_fk_pet_id
        FOREIGN KEY (pet_id) 
        REFERENCES pets (id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB CHARSET=utf8;