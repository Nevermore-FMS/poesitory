CREATE TABLE IF NOT EXISTS upload_tokens (
   id text PRIMARY KEY,
   hash text NOT NULL,
   plugin_id text NOT NULL REFERENCES plugins(id) ON DELETE CASCADE,
   created_at timestamp NOT NULL DEFAULT current_timestamp
);