CREATE TABLE skill (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    level TEXT,
    category TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);
