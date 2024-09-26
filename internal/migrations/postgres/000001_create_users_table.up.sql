CREATE TABLE IF NOT EXISTS users (
    user_id serial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    surname VARCHAR (50) NOT NULL,
    hashed_pass VARCHAR (300) NOT NULL,
    nickname VARCHAR (50) UNIQUE NOT NULL,
    date_reg VARCHAR (50) UNIQUE NOT NULL,
    status VARCHAR (300) NOT NULL
);
