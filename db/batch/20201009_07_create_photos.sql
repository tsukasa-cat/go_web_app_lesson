create table photos (
    id INT(64) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    url TEXT NOT NULL,
    pet_id INT(64) NOT NULL,
    CONSTRAINT photos_fk_pet_id
        FOREIGN KEY (pet_id) 
        REFERENCES pets (id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB CHARSET=utf8;