CREATE TABLE work_experience (
    id SERIAL PRIMARY KEY,
    company TEXT NOT NULL,
    position TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW()
);
