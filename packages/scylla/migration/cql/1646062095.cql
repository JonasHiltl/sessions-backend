CREATE TABLE IF NOT EXISTS story (
  id text PRIMARY KEY,
  party_id text,
  user_id text,
  geohash text,
  url text,
  tagged_friends set<text>,
  created_at timestamp,
) WITH default_time_to_live = 86400;

CREATE INDEX story_by_user ON sessions.story (user_id);
CREATE INDEX story_by_party ON sessions.story (party_id);
