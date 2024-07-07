CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    passport_number VARCHAR(10) NOT NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL,
    address TEXT NOT NULL
)