CREATE TABLE IF NOT EXISTS user_group_menus (
    id SERIAL PRIMARY KEY,
    user_group_id INT NOT NULL,
    menu_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_group_id) REFERENCES user_groups (id) ON DELETE CASCADE,
    FOREIGN KEY (menu_id) REFERENCES menus (id) ON DELETE CASCADE,
    CONSTRAINT unique_user_group_menu UNIQUE (user_group_id, menu_id)
);
