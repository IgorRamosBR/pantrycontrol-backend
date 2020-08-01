CREATE TABLE IF NOT EXISTS products(
   id serial PRIMARY KEY,
   name TEXT NOT NULL,
   unit TEXT,
   brand TEXT,
   category TEXT
);