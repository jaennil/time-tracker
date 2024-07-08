ALTER TABLE users DROP COLUMN passport_serie,
    DROP COLUMN passport_number;

ALTER TABLE users ADD COLUMN passport_number VARCHAR(10) NOT NULL;
