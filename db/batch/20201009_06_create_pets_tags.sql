create table pets_tags (
    pet_id INT(64) NOT NULL,
    tag_id INT(64) NOT NULL,
    CONSTRAINT pets_tags_fk_pet_id
        FOREIGN KEY (pet_id) 
        REFERENCES pets (id)
        ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT pets_tags_fk_tag_id
        FOREIGN KEY (tag_id) 
        REFERENCES tags (id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB CHARSET=utf8;