CREATE TABLE IF NOT EXISTS showcase (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    summary TEXT NOT NULL,
    stack TEXT[] NOT NULL,
    repo_url TEXT NOT NULL,
    live_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
