CREATE TABLE IF NOT EXISTS user_tokens (
   hash text PRIMARY KEY,
   user_id text NOT NULL REFERENCES users(id) ON DELETE CASCADE,
   expires_at timestamp NOT NULL
);