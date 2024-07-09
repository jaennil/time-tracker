CREATE TABLE tasks (
    task_id    SERIAL PRIMARY KEY,
    user_id    INT REFERENCES users (user_id),
    name       VARCHAR(255) NOT NULL,
    start_time TIMESTAMPTZ  NOT NULL,
    end_time   TIMESTAMPTZ
);