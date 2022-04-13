CREATE TABLE IF NOT EXISTS stories (
  id timeuuid PRIMARY KEY,
  party_id text,
  user_id text,
  geohash text,
  url text,
  tagged_friends list<text>,
) WITH default_time_to_live = 86400;

CREATE MATERIALIZED VIEW stories_by_party AS 
  SELECT * FROM stories
  WHERE id IS NOT NULL AND party_id IS NOT NULL
  PRIMARY KEY (party_id, id)
WITH CLUSTERING ORDER BY (id DESC);

CREATE MATERIALIZED VIEW stories_by_user AS 
  SELECT * FROM stories
  WHERE id IS NOT NULL AND user_id IS NOT NULL
  PRIMARY KEY (user_id, id)
WITH CLUSTERING ORDER BY (id DESC);

