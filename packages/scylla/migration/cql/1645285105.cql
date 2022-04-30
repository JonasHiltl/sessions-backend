CREATE TABLE IF NOT EXISTS parties (
  id timeuuid PRIMARY KEY,
  user_id text,
  title text,
  is_public boolean,
  geohash text,
  position list<float>,
  street_address text,
  postal_code text,
  state text,
  country text,
  start_date timestamp,
) WITH default_time_to_live = 86400
AND cdc = {'enabled': true};

CREATE INDEX parties_by_user_test ON sessions.parties (user_id);

CREATE MATERIALIZED VIEW parties_by_user AS 
  SELECT * FROM parties
  WHERE id IS NOT NULL AND user_id IS NOT NULL
  PRIMARY KEY (user_id, id)
WITH CLUSTERING ORDER BY (id DESC);
