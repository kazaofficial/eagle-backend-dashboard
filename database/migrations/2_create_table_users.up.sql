CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    group_id INT NOT NULL REFERENCES user_groups(id) ON DELETE SET NULL ON UPDATE CASCADE,
    name VARCHAR(100) NOT NULL,
    usename VARCHAR(100) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    nrp VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);
