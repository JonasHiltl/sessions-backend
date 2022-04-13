CREATE TABLE IF NOT EXISTS friend_relations (
  user_id timeuuid,
  friend_id timeuuid,
  accepted boolean,
  created_at timestamp,
  accepted_at timestamp,
  PRIMARY KEY (user_id, friend_id)
);

CREATE MATERIALIZED VIEW filtered_friend_relations AS 
  SELECT * FROM friend_relation
  WHERE user_id IS NOT NULL AND friend_id IS NOT NULL AND accepted IS NOT NULL
  PRIMARY KEY (user_id, accepted, created_at)
WITH CLUSTERING ORDER BY (created_at DESC);