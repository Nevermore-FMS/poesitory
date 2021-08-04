CREATE TABLE IF NOT EXISTS plugins (
   id text PRIMARY KEY,
   name character varying (50) UNIQUE NOT NULL,
   type text NOT NULL,
   owner text REFERENCES users(id) NOT NULL
);