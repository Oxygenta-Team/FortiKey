CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(32) NOT NULL UNIQUE CHECK ( login ~ '^[a-zA-Z0-9_]+$' ), --TODO:IN FUTURE FIX IT
    email VARCHAR(255) UNIQUE CHECK(email ~ '^[^@]+@[^@]+\.[a-zA-Z]{2,}$'),
    created_at TIMESTAMP DEFAULT NOW()
)