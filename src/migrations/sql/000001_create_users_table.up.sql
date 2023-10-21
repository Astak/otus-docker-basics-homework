CREATE TABLE IF NOT EXISTS users(
   id bigserial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   firstname VARCHAR (50) NOT NULL,
   lastname VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   phone VARCHAR (12) UNIQUE NOT NULL
);