ALTER TABLE work_experience
ALTER COLUMN description
TYPE TEXT[]
USING
    CASE
        WHEN description IS NULL THEN NULL
        ELSE ARRAY[description]
END;
