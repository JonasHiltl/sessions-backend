CREATE TABLE IF NOT EXISTS comments (
  comment_id timeuuid PRIMARY KEY,
  party_id text,
  author_id text,
  body text,
);

CREATE INDEX comments_by_party_test ON sessions.comments (party_id);
CREATE INDEX comments_by_user_test ON sessions.comments (author_id);

CREATE MATERIALIZED VIEW comments_by_party AS 
  SELECT * FROM comments
  WHERE party_id IS NOT NULL AND comment_id IS NOT NULL
  PRIMARY KEY (party_id, comment_id)
WITH CLUSTERING ORDER BY (comment_id DESC);

CREATE MATERIALIZED VIEW comments_by_user AS 
  SELECT * FROM comments
  WHERE author_id IS NOT NULL AND comment_id IS NOT NULL
  PRIMARY KEY (author_id, comment_id)
WITH CLUSTERING ORDER BY (comment_id DESC);

CREATE TABLE IF NOT EXISTS replies (
  comment_id text,
  author_id text,
  body text,
  created_at timestamp,
  PRIMARY KEY (comment_id, created_at)
) WITH CLUSTERING ORDER BY (created_at ASC);
