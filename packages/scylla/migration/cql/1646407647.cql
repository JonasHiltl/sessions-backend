CREATE TABLE IF NOT EXISTS comments (
  id timeuuid PRIMARY KEY,
  party_id text,
  author_id text,
  body text,
);

CREATE INDEX comments_by_party_test ON sessions.comments (party_id);
CREATE INDEX comments_by_user_test ON sessions.comments (author_id);

CREATE MATERIALIZED VIEW comments_by_party AS 
  SELECT * FROM comments
  WHERE party_id IS NOT NULL AND id IS NOT NULL
  PRIMARY KEY (party_id, id)
WITH CLUSTERING ORDER BY (id DESC);

CREATE MATERIALIZED VIEW comments_by_user AS 
  SELECT * FROM comments
  WHERE author_id IS NOT NULL AND id IS NOT NULL
  PRIMARY KEY (author_id, id)
WITH CLUSTERING ORDER BY (id DESC);

CREATE TABLE IF NOT EXISTS replies (
  id text,
  author_id text,
  body text,
  created_at timestamp,
  PRIMARY KEY (id, created_at)
) WITH CLUSTERING ORDER BY (created_at ASC);
