CREATE TABLE IF NOT EXISTS user_group_menus (
    id SERIAL PRIMARY KEY,
    user_group_id INT NOT NULL,
    menu_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_group_id) REFERENCES user_groups (id),
    FOREIGN KEY (menu_id) REFERENCES menus (id)
);
