CREATE TABLE IF NOT EXISTS comments (
  comment_id text,
  party_id text,
  author_id text,
  body text,
  created_at timestamp,
  PRIMARY KEY (comment_id, party_id)
);

CREATE MATERIALIZED VIEW commenty_by_party AS 
  SELECT * FROM comments
  WHERE party_id IS NOT NULL AND comment_id IS NOT NULL AND author_id IS NOT NULL AND created_at IS NOT NULL
  PRIMARY KEY (party_id, created_at, comment_id)
WITH CLUSTERING ORDER BY (created_at DESC);

CREATE TABLE IF NOT EXISTS replies (
  comment_id text,
  author_id text,
  body text,
  created_at timestamp,
  PRIMARY KEY (comment_id, created_at)
) WITH CLUSTERING ORDER BY (created_at ASC);
