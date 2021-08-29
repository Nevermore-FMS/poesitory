CREATE TABLE IF NOT EXISTS plugin_versions (
   id text PRIMARY KEY,
   plugin text NOT NULL REFERENCES plugins(id) ON DELETE CASCADE,
   hash text NOT NULL,
   major integer NOT NULL,
   minor integer NOT NULL,
   patch integer NOT NULL,
   channel character varying(50),
   timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   readme text,
   UNIQUE(plugin, major, minor, patch, channel)
);