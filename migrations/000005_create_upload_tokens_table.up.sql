CREATE TABLE IF NOT EXISTS upload_tokens (
   hash text PRIMARY KEY,
   plugin_id text NOT NULL REFERENCES plugins(id) ON DELETE CASCADE,
   created_at timestamp NOT NULL DEFAULT current_timestamp
);