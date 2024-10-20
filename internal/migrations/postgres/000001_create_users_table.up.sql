CREATE TABLE IF NOT EXISTS users (
    user_id serial PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    hashed_pass TEXT NOT NULL,
    nickname TEXT UNIQUE NOT NULL,
    date_reg timestamp with timezone NOT NULL default now(),
    status TEXT
);
