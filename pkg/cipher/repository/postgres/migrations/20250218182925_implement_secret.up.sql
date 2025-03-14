CREATE TABLE IF NOT EXISTS secrets(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    method VARCHAR(45) NOT NULL,
    key VARCHAR(255) NOT NULL UNIQUE,
    hash BYTEA NOT NULL,
    is_deleted BOOLEAN DEFAULT false
);