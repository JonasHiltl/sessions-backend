CREATE TABLE IF NOT EXISTS story (
  id text,
  party_id text,
  user_id text,
  geohash text,
  url text,
  tagged_friends list<text>,
  created_at timestamp,
  PRIMARY KEY (id, party_id)
) WITH default_time_to_live = 86400;

CREATE INDEX story_by_user ON sessions.story (user_id);

CREATE MATERIALIZED VIEW story_by_party AS 
  SELECT * FROM story
  WHERE id IS NOT NULL AND party_id IS NOT NULL AND user_id IS NOT NULL AND created_at IS NOT NULL
  PRIMARY KEY (party_id, created_at, id)
  WITH comment='Allow query by party and sorted by creation date';


