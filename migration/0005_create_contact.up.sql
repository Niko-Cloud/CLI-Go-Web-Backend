CREATE TABLE contact (
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);
