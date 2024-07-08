ALTER TABLE users DROP COLUMN passport_number;

ALTER TABLE users ADD COLUMN passport_serie VARCHAR(4) NOT NULL,
    ADD COLUMN passport_number VARCHAR(6) NOT NULL;