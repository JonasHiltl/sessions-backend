CREATE TABLE IF NOT EXISTS party (
  id timeuuid PRIMARY KEY,
  user_id text,
  title text,
  is_public boolean,
  geohash text,
  position list<float>,
  start_date timestamp,
) WITH default_time_to_live = 86400;

CREATE INDEX party_by_user_test ON sessions.party (user_id);

CREATE MATERIALIZED VIEW party_by_user AS 
  SELECT * FROM party
  WHERE id IS NOT NULL AND user_id IS NOT NULL
  PRIMARY KEY (user_id, id)
WITH CLUSTERING ORDER BY (id DESC);
