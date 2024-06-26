CREATE TABLE IF NOT EXISTS menus (
  id SERIAL PRIMARY KEY,
  parent_id INT REFERENCES menus(id) ON DELETE SET NULL ON UPDATE CASCADE,
  url_key VARCHAR(255) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  icon VARCHAR(255),
  url VARCHAR(255),
  is_shown BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
