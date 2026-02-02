ALTER TABLE skill
ALTER COLUMN category
TYPE TEXT
USING
    CASE
        WHEN category IS NULL THEN NULL
        ELSE description[1]
END;
