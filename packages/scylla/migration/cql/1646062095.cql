CREATE TABLE IF NOT EXISTS story (
  id timeuuid PRIMARY KEY,
  party_id text,
  user_id text,
  geohash text,
  url text,
  tagged_friends list<text>,
) WITH default_time_to_live = 86400;

CREATE INDEX story_by_party_test ON sessions.story (party_id);
CREATE INDEX story_by_user_test ON sessions.story (user_id);

CREATE MATERIALIZED VIEW story_by_party AS 
  SELECT * FROM story
  WHERE id IS NOT NULL AND party_id IS NOT NULL
  PRIMARY KEY (party_id, id)
WITH CLUSTERING ORDER BY (id DESC);

CREATE MATERIALIZED VIEW story_by_user AS 
  SELECT * FROM story
  WHERE id IS NOT NULL AND user_id IS NOT NULL
  PRIMARY KEY (user_id, id)
WITH CLUSTERING ORDER BY (id DESC);

