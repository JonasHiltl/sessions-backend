CREATE TABLE IF NOT EXISTS party (
  id text PRIMARY KEY,
  user_id text,
  title text,
  is_public boolean,
  geohash text,
  created_at timestamp,
) WITH default_time_to_live = 86400;

CREATE INDEX party_by_user ON sessions.party (user_id);