CREATE TABLE IF NOT EXISTS plugin_versions (
   id text PRIMARY KEY,
   plugin text REFERENCES plugins(id) NOT NULL,
   hash text UNIQUE NOT NULL,
   major integer NOT NULL,
   minor integer NOT NULL,
   patch integer NOT NULL,
   channel character varying(50),
   timestamp timestamp,
   readme text
);