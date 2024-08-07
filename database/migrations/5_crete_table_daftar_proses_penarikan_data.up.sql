CREATE TABLE IF NOT EXISTS daftar_proses_penarikan_data (
    id SERIAL PRIMARY KEY,
    conn_type VARCHAR NOT NULL,
    source_connection_id VARCHAR NOT NULL,
    target_connection_id VARCHAR NOT NULL,
    schema_table VARCHAR NOT NULL,
    schedule VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
