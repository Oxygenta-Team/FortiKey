CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    login VARCHAR(32) NOT NULL UNIQUE CHECK ( login ~ '^[a-zA-Z0-9_]+$' ),
    email VARCHAR(255) UNIQUE CHECK(email ~ '^[^@]+@[^@]+\.[a-zA-Z]{2,}$'),
    is_delete BOOLEAN DEFAULT FALSE;
)